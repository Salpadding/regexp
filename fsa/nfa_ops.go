package fsa

func (n *NFA) input(r rune) *NFA {
	if n.rejected {
		return n
	}
	if n.currentStates == nil {
		n.currentStates = newStateSet(0)
	}
	t, ok := n.transitions[r]
	if !ok {
		n.rejected = true
		return n
	}
	tmp := n.closure(n.currentStates)
	tmp2 := newStateSet()
	for _, s := range tmp.elements() {
		tmp2 = tmp2.union(t[s])
	}
	n.currentStates = tmp2
	if tmp2.size() == 0 {
		n.rejected = true
	}
	return n
}

func (n *NFA) Input(rs ...rune) FSA {
	for _, r := range rs {
		n.input(r)
	}
	return n
}

func (n *NFA) InputString(s string) FSA {
	for _, r := range s {
		n.Input(r)
	}
	return n
}

func (n *NFA) IsAccept() bool {
	if n.rejected {
		return false
	}
	if n.currentStates == nil {
		n.currentStates = newStateSet(0)
	}
	return n.closure(n.currentStates).intersection(n.finalStates).size() > 0
}

// epsilon closure of states
func (n *NFA) closure(set stateSet) stateSet {
	closure := set.copy()
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

// util method for thompson construction
func (n *NFA) offset(i int) *NFA {
	transitions := make(map[rune]map[state]stateSet, len(n.transitions))
	for r, v := range n.transitions {
		transitions[r] = make(map[state]stateSet, len(v))
		for l, m := range v {
			transitions[r][l+state(i)] = m.offset(i)
		}
	}
	return &NFA{
		transitions:  transitions,
		finalStates:  n.finalStates.offset(i),
		maximumState: n.maximumState + state(i),
	}
}

// util method for thompson construction
func unionTrans(t1 map[state]stateSet, t2 map[state]stateSet) map[state]stateSet {
	res := make(map[state]stateSet)
	for k, v := range t1 {
		_, ok := t2[k]
		if ok {
			panic("transition union fail") // when offset not correctly added
		}
		res[k] = v
	}
	for k, v := range t2 {
		_, ok := t1[k]
		if ok {
			panic("transition union fail")
		}
		res[k] = v
	}
	return res
}

// util method for thompson construction
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

// add a transition rule
func (n *NFA) addTransition(r rune, from, to state) *NFA {
	_, ok := n.transitions[r]
	if !ok {
		n.transitions[r] = make(map[state]stateSet)
	}
	_, ok = n.transitions[r][from]
	if !ok {
		n.transitions[r][from] = newStateSet()
	}
	n.transitions[r][from].add(to)
	return n
}

// thompson construction for ab
// 1. offset the later nfa
// 2. add an epsilon transition between them
func (n *NFA) concat(n1 *NFA) *NFA {
	n2 := n1.offset(int(n.maximumState) + 1)
	res := &NFA{
		finalStates:  n2.finalStates,
		transitions:  unionTransitions(n.transitions, n2.transitions),
		maximumState: n2.maximumState,
	}
	for _, s := range n.finalStates.elements() {
		res.addTransition(epsilon, s, n.maximumState+1)
	}
	return res
}

// thompson construction fo a|b
// 1. offset both nfa
// 2. add an epsilon from 0 to them
// 3. add epsilon transitions to final state
func (n *NFA) or(n1 *NFA) *NFA {
	n3 := n.offset(1)
	n4 := n1.offset(int(n3.maximumState) + 1)
	res := &NFA{
		finalStates:  newStateSet(n4.maximumState + 1),
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

// thompson construction fo a*
// 1. offset the nfa
// 2. add an epsilon from 0 to final state
// 3. add epsilon transitions to final state
// 4. add epsilon transition to 1
func (n *NFA) kleen() *NFA {
	n1 := n.offset(1)
	res := &NFA{
		finalStates:  newStateSet(n1.maximumState + 1),
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
