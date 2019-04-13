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
	states        *states        // states represented by graph
	transitionSet *transitionSet // transation functions of this automata
	current       int            // current state of the nfa, typicallly initialzed 1
	finalStates   *integerSet
	initialState  int
}

type transition struct {
	to   int
	char rune
}

// [1] -- a --> [[2]]
func newNFA(char rune) *nfa {
	var finalStates integerSet = map[int]bool{
		2: true,
	}
	var ts transitions = []*transition{
		&transition{
			to:   2,
			char: char,
		},
	}
	var states states = []int{1, 2}
	var transitionSet transitionSet = map[int]*transitions{
		1: &ts,
	}
	return &nfa{
		states:        &states,
		transitionSet: &transitionSet,
		current:       1,
		finalStates:   &finalStates,
		initialState:  1,
	}
}

func (n *nfa) Input(r rune) {
	if n.current == -1 {
		return
	}
	if !n.transitionSet.has(n.current) {
		n.current = -1
		return
	}
	ts := n.transitionSet.get(n.current)
	for _, t := range *ts {
		if t.char == r {
			n.current = t.to
			return
		}
	}
	n.current = -1
}

func (n *nfa) Reset() {
	n.current = 0
}

func (n *nfa) InputString(s string) {
	for _, r := range s {
		n.Input(r)
	}
}

func (n *nfa) IsAccept() bool {
	return n.finalStates.has(n.current)
}

// [1] -- a --> [[2]] + [1] -- b --> [[2]] ===> [1] -- a --> [2] -- b --> [[3]]
func (n *nfa) concat(n2 *nfa) *nfa {

	// append offsets
	offset := n.states.len() - 1

	return &nfa{
		states: n.states.concat(
			n2.states.slice(1, n2.states.len()).offset(offset),
		),
		transitionSet: n.transitionSet.union(
			n2.transitionSet.offset(offset),
		),
		current:     1,
		finalStates: n2.finalStates.offset(offset),
	}
}

func (n *nfa) or(n2 *nfa) *nfa {
	offset := n.states.len() - 1

	t1 := n2.transitionSet.remove(1)
	t1 = t1.offset(offset)
	transitions := n2.transitionSet.offset(offset)
	transitions.add(1, t1)
	newFinals := n.finalStates.union(
		n2.finalStates.offset(offset),
	)
	return &nfa{
		states: n.states.concat(
			n2.states.slice(1, n2.states.len()).offset(offset),
		),
		transitionSet: n.transitionSet.union(
			transitions,
		),
		current:     1,
		finalStates: newFinals,
	}
}

func (n *nfa) closure() *nfa {
	for _, k := range n.finalStates.entries() {
		n.transitionSet.add(k, n.transitionSet.get(1))
	}
	return n
}
