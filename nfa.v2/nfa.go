package nfa_v2

type state int

const (
	epsilon = 0xffffffff
)

func newStateSet() stateSet {
	return make(map[state]bool)
}

type stateSet map[state]bool

func (set stateSet) has(s state) bool {
	return set[s]
}

func (set stateSet) add(s state) stateSet {
	set[s] = true
	return set
}

func (set stateSet) remove(s state) stateSet {
	set[s] = false
	return set
}

func (set stateSet) union(s2 stateSet) stateSet {
	res := make(map[state]bool)
	for s, ok := range set {
		if ok {
			res[s] = true
		}
	}
	for s, ok := range s2 {
		if ok {
			res[s] = true
		}
	}
	return res
}

func (set stateSet) copy() stateSet {
	res := make(map[state]bool)
	for s, ok := range set {
		if ok {
			res[s] = true
		}
	}
	return res
}

func (set stateSet) size() int {
	i := 0
	for _, ok := range set {
		if ok {
			i++
		}
	}
	return i
}

func (set stateSet) elements() []state {
	res := make([]state, 0)
	for s, ok := range set {
		if ok {
			res = append(res, s)
		}
	}
	return res
}

func (set stateSet) intersection(s2 stateSet) stateSet {
	res := make(map[state]bool, 0)
	for s, ok := range set {
		if !ok {
			continue
		}
		ok = s2.has(s)
		if !ok {
			continue
		}
		res[s] = true
	}
	return res
}

// NFA is non deterministic automaton
type NFA struct {
	// in nfa, we assume the initial state is always zero
	// all possibly state continuously starts from zero

	finalStates stateSet                    // accepted states or final states
	transitions map[rune]map[state]stateSet // transitions, may contains epsilon transitions

	maximumState state // maximum state
}

// epsilon closure of states
func (n *NFA) closure(set stateSet) stateSet {
	closure := stateSet{}.copy()
	fn := func() {
		epsilonTransitions, ok := n.transitions[epsilon]
		if !ok {
			return
		}
		for _, s := range closure.elements() {
			s1, ok := epsilonTransitions[s]
			if !ok {
				continue
			}
			closure = closure.union(s1)
		}
	}
	size0 := closure.size()
	for {
		fn()
		size1 := closure.size()
		if size1 == size0 {
			break
		}
		size0 = size1
	}
	return closure
}

func NewRune(r rune) *NFA {
	return &NFA{
		finalStates: newStateSet().add(1),
		transitions: map[rune]map[state]stateSet{
			r: {0: newStateSet().add(1)},
		},
		maximumState: 1,
	}
}

// util method for thompson construction
func (n *NFA) offset(i int) *NFA {
	finalStates := make(map[state]bool, len(n.finalStates))
	for k, ok := range n.finalStates {
		if !ok {
			continue
		}
		finalStates[k+state(i)] = true
	}
	transitions := make(map[rune]map[state]stateSet, len(n.transitions))
	for r, v := range n.transitions {
		transitions[r] = make(map[state]stateSet, len(v))
		for l, m := range v {
			m2 := newStateSet()
			for _, s := range m.elements() {
				m2.add(s + state(i))
			}
			transitions[r][l+state(i)] = m2
		}
	}

	return &NFA{
		transitions:  transitions,
		finalStates:  finalStates,
		maximumState: n.maximumState + state(i),
	}
}

func unionTrans(t1 map[state]stateSet, t2 map[state]stateSet) map[state]stateSet {
	res := make(map[state]stateSet)
	for k, v := range t1 {
		res[k] = v
	}
	for k, v := range t2 {
		res[k] = v
	}
	return res
}

func unionTransitions(tss1 map[rune]map[state]stateSet, tss2 map[rune]map[state]stateSet) map[rune]map[state]stateSet {
	res := make(map[rune]map[state]stateSet)
	for r, ts := range tss1 {
		_, ok := res[r]
		if !ok {
			res[r] = make(map[state]stateSet)
		}
		res[r] = unionTrans(ts, res[r])
	}
	for r, ts := range tss2 {
		_, ok := res[r]
		if !ok {
			res[r] = make(map[state]stateSet)
		}
		res[r] = unionTrans(ts, res[r])
	}
	return res
}

func (n *NFA) addTransition(r rune, from, to state) {
	_, ok := n.transitions[r]
	if !ok {
		n.transitions[r] = make(map[state]stateSet)
	}
	_, ok = n.transitions[r][from]
	if !ok {
		n.transitions[r][from] = newStateSet()
	}
	n.transitions[r][from].add(to)
}

// thompson construction
func (n *NFA) concat(n1 *NFA) *NFA {
	n2 := n1.offset(int(n.maximumState) + 1)
	res := &NFA{
		finalStates:  n2.finalStates,
		transitions:  unionTransitions(n.transitions, n2.transitions),
		maximumState: n1.maximumState,
	}
	for _, s := range n.finalStates.elements() {
		res.addTransition(epsilon, s, n.maximumState+1)
	}
	return res
}

// thompson construction
func (n *NFA) or(n1 *NFA) *NFA {
	n3 := n.offset(1)
	n4 := n1.offset(int(n3.maximumState) + 1)
	res := &NFA{
		finalStates:  newStateSet().add(n4.maximumState + 1),
		transitions:  unionTransitions(n3.transitions, n4.transitions),
		maximumState: n4.maximumState + 1,
	}
	res.addTransition(epsilon, 0, 1)
	res.addTransition(epsilon, 0, n3.maximumState+1)
	for _, s := range n3.finalStates.elements() {
		res.addTransition(epsilon, s, n4.maximumState+1)
	}
	for _, s := range n4.finalStates.elements() {
		res.addTransition(epsilon, s, n4.maximumState+1)
	}
	return res
}

func (n *NFA) kleen() *NFA {
	n1 := n.offset(1)
	res := &NFA{
		finalStates:  newStateSet().add(n1.maximumState + 1),
		transitions:  n1.transitions,
		maximumState: n1.maximumState + 1,
	}
	res.addTransition(epsilon, 0, 1)
	res.addTransition(epsilon, 0, n1.maximumState+1)
	for _, s := range n1.finalStates.elements() {
		res.addTransition(epsilon, s, 1)
		res.addTransition(epsilon, s, n1.maximumState+1)
	}
	return res
}
