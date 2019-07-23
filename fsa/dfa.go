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

// converting an NFA to a DFA
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
		transitions: trans,
		finalStates: newStateSet(),
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
}

func (d *DFA) Input(rs ...rune) {
	if d.rejected {
		return
	}
	for _, r := range rs {
		_, ok := d.transitions[r]
		if !ok {
			d.rejected = true
			return
		}
		s2, ok := d.transitions[r][d.currentState]
		if !ok {
			d.rejected = true
			return
		}
		d.currentState = s2
	}
}

func (d *DFA) InputString(s string) {
	d.Input([]rune(s)...)
}

func (d *DFA) IsAccept() bool {
	return !d.rejected && d.finalStates.has(d.currentState)
}
