package fsa

func (n *NFA) oneOrMore() *NFA {
	return n.concat(n.kleen())
}

func (n *NFA) noneOrOne() *NFA {
	return n.or(newEpsilon())
}

func newEpsilon() *NFA {
	return &NFA{
		finalStates: newStateSet(1),
		transitions: map[rune]map[state]stateSet{
			epsilon: {0: newStateSet(1)},
		},
		maximumState: 1,
	}
}
