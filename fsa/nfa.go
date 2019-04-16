package fsa

var epsilon rune = 0

type nfa struct {
	states      []int
	f           TransitionSet
	finalStates *integerSet
}

func (n *nfa) toDFA() *nfa {
	for _, s := range n.states {
		moves := n.f.epsilonMoves(s)
	}
}
