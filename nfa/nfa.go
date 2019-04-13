package nfa

const (
	leftParentheses       = '('
	rightParentheses      = ')'
	escape                = '\\'
	or                    = '|'
	closure               = '*'
	whiteSpace            = ' '
	epsilon          rune = 0
)

type nfa struct {
	states       []int                 // states represented by graph
	transitions  map[int][]*transition // transation functions of this automata
	current      int                   // current state of the nfa, typicallly initialzed 1
	finalStates  []int
	initialState int
}

type transition struct {
	to   int
	char rune
}

// [1] -- a --> [[2]]
func newNFA(char rune) *nfa {
	return &nfa{
		states: []int{1, 2},
		transitions: map[int][]*transition{
			1: []*transition{
				&transition{
					to:   2,
					char: char,
				},
			},
		},
		current:      1,
		finalStates:  []int{2},
		initialState: 1,
	}
}

func (n *nfa) Input(r rune) {
	if n.current == -1 {
		return
	}
	ts := n.transitions[n.current]
	for _, t := range ts {
		if t.char == r {
			n.current = t.to
			return
		}
	}
	n.current = -1
}

func (n *nfa) InputString(s string) {
	for _, r := range s {
		n.Input(r)
	}
}

func (n *nfa) IsAccept() bool {
	for _, s := range n.finalStates {
		if s == n.current {
			return true
		}
	}
	return false
}

// [1] -- a --> [[2]] + [1] -- b --> [[2]] ===> [1] -- a --> [2] -- b --> [[3]]
func (n *nfa) concat(n2 *nfa) *nfa {

	// append offsets
	offset := len(n.states) - 1
	newStates := make([]int, len(n.states)+len(n2.states)-1)
	n2StatesPadded := make([]int, len(n2.states))
	for idx, s := range n2.states {
		n2StatesPadded[idx] = offset + s
	}
	copy(newStates, append(n.states, n2StatesPadded[1:]...))

	newFinals := make([]int, len(n2.finalStates))
	for idx, s := range n2.finalStates {
		newFinals[idx] = offset + s
	}
	transitions := make(map[int][]*transition, len(n.transitions)+len(n2.transitions))
	for k, v := range n2.transitions {
		_, ok := transitions[k]
		if !ok {
			transitions[k+offset] = make([]*transition, len(v))
		}
		for idx, t := range v {
			transitions[k+offset][idx] = &transition{
				to:   t.to + offset,
				char: t.char,
			}
		}
	}
	for k, v := range n.transitions {
		_, ok := transitions[k]
		if !ok {
			transitions[k] = make([]*transition, len(v))
		}
		for idx, t := range v {
			transitions[k][idx] = &transition{
				to:   t.to,
				char: t.char,
			}
		}
	}
	return &nfa{
		states:      newStates,
		transitions: transitions,
		current:     1,
		finalStates: newFinals,
	}
}
