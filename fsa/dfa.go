package fsa

func (n *NFA) trans(alpha rune, in stateSet) stateSet {
	tmp := newStateSet()
	t, ok := n.transitions[alpha]
	if !ok {
		return tmp
	}
	for _, s := range in.elements() {
		s1, ok := t[s]
		if ok {
			tmp = tmp.union(s1)
		}
	}
	return n.closure(tmp)
}

// converting an NFA to a DFA, implements http://web.cecs.pdx.edu/~harry/compilers/slides/LexicalPart3.pdf
func (n *NFA) ToDFA() *DFA {
	// each state in the DFA will correspond to a set of NFA states.
	dfaStates := []stateSet{n.closure(map[state]bool{0: true})}
	traversed := make(map[int]bool)
	trans := make(map[rune]map[state]state)

	indexOf := func(set stateSet) int {
		for i, s := range dfaStates {
			if set.equal(s) {
				return i
			}
		}
		dfaStates = append(dfaStates, set)
		return len(dfaStates) - 1
	}

	addTrans := func(alpha rune, from, to stateSet) {
		_, ok := trans[alpha]
		if !ok {
			trans[alpha] = make(map[state]state)
		}
		i1, i2 := indexOf(from), indexOf(to)
		trans[alpha][state(i1)] = state(i2)
	}

	fn := func() {
		tmp := make([]stateSet, len(dfaStates))
		copy(tmp, dfaStates)
		for i, s := range tmp {
			if traversed[i] {
				continue
			}
			traversed[i] = true
			for alpha := range n.transitions {
				if alpha == epsilon {
					continue
				}
				res := n.trans(alpha, s)
				addTrans(alpha, s, res)
			}
		}
	}

	size0 := len(dfaStates)
	for {
		fn()
		size1 := len(dfaStates)
		if size1 == size0 {
			break
		}
		size0 = size1
	}

	dfa := &DFA{
		transitions:  trans,
		finalStates:  newStateSet(),
		maximumState: state(len(dfaStates) - 1),
	}

	for i, s := range dfaStates {
		if s.intersection(n.finalStates).size() > 0 {
			dfa.finalStates.add(state(i))
		}
	}

	return dfa
}

type DFA struct {
	currentState state // a deterministic automaton contains one current state

	// the result of transition from a state is deterministic
	// there exists none epsilon transitions
	transitions map[rune]map[state]state

	finalStates stateSet

	rejected bool

	maximumState state
}

// dfa is a special case of nfa
func (d *DFA) ToNFA() *NFA {
	nfa := &NFA{
		finalStates:  d.finalStates,
		transitions:  make(map[rune]map[state]stateSet, len(d.transitions)),
		maximumState: d.maximumState,
	}
	for k, v := range d.transitions {
		for from, to := range v {
			nfa.addTransition(k, from, to)
		}
	}
	return nfa
}

func (d *DFA) Input(rs ...rune) FSA {
	if d.rejected {
		return d
	}
	for _, r := range rs {
		_, ok := d.transitions[r]
		if !ok {
			d.rejected = true
			return d
		}
		s2, ok := d.transitions[r][d.currentState]
		if !ok {
			d.rejected = true
			return d
		}
		d.currentState = s2
	}
	return d
}

func (d *DFA) InputString(s string) FSA {
	d.Input([]rune(s)...)
	return d
}

func (d *DFA) IsAccept() bool {
	return !d.rejected && d.finalStates.has(d.currentState)
}

func (d *DFA) Reset() FSA {
	d.currentState = 0
	d.rejected = false
	return d
}

func (d *DFA) closure(set stateSet) stateSet {
	res := set.copy()

	fn := func() {
		for _, s := range res.elements() {
			for alpha := range d.transitions {
				s1, ok := d.transitions[alpha][s]
				if ok {
					res.add(s1)
				}
			}
		}
	}

	size0 := res.size()
	for {
		fn()
		size1 := res.size()
		if size1 == size0 {
			break
		}
		size0 = size1
	}
	return res
}

// hopcroft dfa Minimize
func (d *DFA) Minimize() *DFA {
	sts := d.refineAll()
	for i := 1; i < len(sts); i++ {
		if sts[i].has(0) {
			tmp := sts[i]
			sts[i] = sts[0]
			sts[0] = tmp
			break
		}
	}
	// filter dead state
	var tmp []stateSet
	for _, set := range sts {
		if d.closure(set).intersection(d.finalStates).size() != 0 {
			tmp = append(tmp, set)
		}
	}
	sts = tmp
	res := &DFA{
		transitions:  map[rune]map[state]state{},
		maximumState: state(len(sts) - 1),
		finalStates:  newStateSet(),
	}

	indexOf := func(s state) int {
		for i := range sts {
			if sts[i].has(s) {
				return i
			}
		}
		return -1
	}

	addTransition := func(alpha rune, from state, to state) {
		_, ok := res.transitions[alpha]
		if !ok {
			res.transitions[alpha] = map[state]state{}
		}
		res.transitions[alpha][from] = to
	}

	for i := range sts {
		s, ok := sts[i].one()
		if !ok {
			continue
		}
		if d.finalStates.has(s) {
			res.finalStates.add(state(i))
		}
		for alpha := range d.transitions {
			s2, ok := d.transitions[alpha][s]
			if !ok {
				continue
			}
			to := indexOf(s2)
			if to < 0 {
				continue
			}
			addTransition(alpha, state(i), state(to))
		}
	}
	return res
}

// hopcroft's algorithm partition refinement
func (d *DFA) refineAll() []stateSet {
	sts := newStateSet()
	for i := state(0); i <= d.maximumState; i++ {
		sts.add(i)
	}
	for _, s := range d.finalStates.elements() {
		sts.remove(s)
	}
	initial := []stateSet{sts, d.finalStates}
	size0 := len(initial)
	for {
		for alpha := range d.transitions {
			initial = d.refineOneChar(initial, alpha)
		}
		size1 := len(initial)
		if size1 == size0 {
			break
		}
		size0 = size1
	}
	return initial
}

// hopcroft's algorithm partition refinement
func (d *DFA) refineOneChar(sts []stateSet, alpha rune) []stateSet {
	var res []stateSet
	ts := d.transitions[alpha]
	indexOf := func(s state) int {
		for i := range sts {
			if sts[i].has(s) {
				return i
			}
		}
		return -1
	}

	for _, set := range sts {
		m := make(map[state]stateSet)
		for _, s := range set.elements() {
			t, ok := ts[s]
			if !ok {
				t = -1
			}
			_, ok = m[t]
			if !ok {
				m[t] = newStateSet()
			}
			m[t].add(s)
		}

		m2 := make(map[int]stateSet)
		for k, v := range m {
			idx := indexOf(k)
			_, ok := m2[idx]
			if !ok {
				m2[idx] = newStateSet()
			}
			m2[idx] = m2[idx].union(v)
		}

		for _, v := range m2 {
			res = append(res, v)
		}
	}
	return res
}
