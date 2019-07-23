package fsa

const (
	// epsilon transition
	// 0xffff is large enough to avoid char code collision
	epsilon = 0xffff
)

// NFA is non deterministic automaton
type NFA struct {
	// in nfa, we assume the initial state is always zero
	// all possibly state continuously starts from zero

	finalStates stateSet                    // accepted states or final states
	transitions map[rune]map[state]stateSet // contains epsilon transitions, the transition result of an input is non deterministic (a set of results)

	maximumState state // maximum state number, for quickly thompson construction

	currentStates stateSet
	rejected      bool // whether in error state
}

func NewChar(r rune) *NFA {
	n := newEmpty()
	return n.addTransition(r, 0, 1)
}

func newEmpty() *NFA {
	return &NFA{
		finalStates: newStateSet(1),
		transitions: map[rune]map[state]stateSet{
		},
		maximumState: 1,
	}
}

func (n *NFA) oneOrMore() *NFA {
	return n.concat(n.kleen())
}

func (n *NFA) noneOrOne() *NFA {
	return n.or(newEpsilon())
}

func newEpsilon() *NFA {
	return newEmpty().addTransition(epsilon, 0, 1)
}

// new wildcard nfa
func newWildCard() *NFA {
	res := newEmpty()
	for r := rune(0); r < 0x7f; r++ {
		res.addTransition(r, 0, 1)
	}
	return res
}

func newDigital() *NFA {
	res := newEmpty()
	for r := '0'; r <= '9'; r++ {
		res.addTransition(r, 0, 1)
	}
	return res
}

func newLetters() *NFA {
	res := newEmpty()
	for r := 'a'; r <= 'z'; r++ {
		res.addTransition(r, 0, 1)
	}
	for r := 'A'; r <= 'Z'; r++ {
		res.addTransition(r, 0, 1)
	}
	return res
}

func newRange(start, end rune) *NFA {
	res := newEmpty()
	for r := start; r <= end; r++ {
		res.addTransition(r, 0, 1)
	}
	return res
}

func newNonDigital() *NFA {
	res := newEmpty()
	for r := rune(0); r < 0x7f && !('0' <= r && r <= '9'); r++ {
		res.addTransition(r, 0, 1)
	}
	return res
}

func newNonLetter() *NFA {
	res := newEmpty()
	for r := rune(0); r < 0x7f && !('a' <= r && r <= 'z') && !('A' <= r && r <= 'Z'); r++ {
		res.addTransition(r, 0, 1)
	}
	return res
}

func (n *NFA) Reset() RegExp {
	n.currentStates = newStateSet(0)
	return n
}
