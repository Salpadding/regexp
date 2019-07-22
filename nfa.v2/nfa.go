package nfa_v2

type state int

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

type NFA struct {
	currentStates      stateSet
	finalStates        stateSet
	transitions        map[rune]map[state]state
	epsilonTransitions map[state]state
	rejected           bool
}

func (n *NFA) Input(r rune) {
	if n.rejected {
		return
	}
	t, ok := n.transitions[r]
	if !ok {
		n.rejected = true
		return
	}
	for _, s := range n.currentStates.elements() {
		n.currentStates.remove(s)
		s1, ok := t[s]
		if !ok {
			continue
		}
		n.currentStates.add(s1)
	}
	if n.currentStates.size() == 0 {
		n.rejected = true
		return
	}
}

func (n *NFA) Accepted() bool {
	if n.rejected {
		return false
	}
	closure := n.currentStates.copy()
	fn := func() {
		for _, s := range closure.elements() {
			s1, ok := n.epsilonTransitions[s]
			if !ok {
				continue
			}
			closure.add(s1)
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
	return closure.intersection(n.finalStates).size() > 0
}

func NewRune(r rune) *NFA {
	return &NFA{
		currentStates: newStateSet().add(0),
		finalStates:   newStateSet().add(1),
		transitions: map[rune]map[state]state{
			r: {0: 1},
		},
		epsilonTransitions: make(map[state]state),
	}
}
