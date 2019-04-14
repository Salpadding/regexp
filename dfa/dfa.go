package dfa

const (
	leftParentheses       = '('
	rightParentheses      = ')'
	escape                = '\\'
	or                    = '|'
	closure               = '*'
	whiteSpace            = ' '
	dot                   = '.'
	epsilon          rune = 0
)

type dfa struct {
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
func newNFA(char rune) *dfa {
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
	return &dfa{
		states:        &states,
		transitionSet: &transitionSet,
		current:       1,
		finalStates:   &finalStates,
		initialState:  1,
	}
}

func (n *dfa) Input(r rune) {
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

func (n *dfa) Reset() {
	n.current = 0
}

func (n *dfa) InputString(s string) {
	for _, r := range s {
		n.Input(r)
	}
}

func (n *dfa) IsAccept() bool {
	return n.finalStates.has(n.current)
}

// [1] -- a --> [[2]] + [1] -- b --> [[2]] ===> [1] -- a --> [2] -- b --> [[3]]
func (n *dfa) concat(n2 *dfa) *dfa {
	// 核心思想是增加把 1 的 final states 作为 2 的初始状态

	// append offsets
	offset := n.states.len() - 1
	copied2 := n2.transitionSet.offset(0)
	t1 := copied2.remove(1)
	t1 = t1.offset(offset)
	copied1 := n.transitionSet.offset(0)

	for _, k := range n.finalStates.entries() {
		copied1.add(k, t1)
	}
	return &dfa{
		states: n.states.concat(
			n2.states.slice(1, n2.states.len()).offset(offset),
		),
		transitionSet: copied1.union(copied2.offset(offset)),
		current:       1,
		finalStates:   n2.finalStates.offset(offset),
	}
}

func (n *dfa) or(n2 *dfa) *dfa {
	// 核心思想是取两个final state 为并集, 把 2 的初始状态对应的 trans 加上 offset 给1
	offset := n.states.len() - 1

	copied := n2.transitionSet.copy()
	t1 := copied.remove(1)
	t1 = t1.offset(offset)
	copied = copied.offset(offset).add(1, t1)
	newFinals := n.finalStates.union(
		n2.finalStates.offset(offset),
	)
	return &dfa{
		states: n.states.concat(
			n2.states.slice(1, n2.states.len()).offset(offset),
		),
		transitionSet: n.transitionSet.union(
			copied,
		),
		current:     1,
		finalStates: newFinals,
	}
}

func (n *dfa) closure() *dfa {
	copied := n.transitionSet.offset(0)
	for _, k := range n.finalStates.entries() {
		copied = copied.add(k, n.transitionSet.get(1))
	}
	fcopied := n.finalStates.offset(0)
	fcopied.add(1)
	return &dfa{
		states:        n.states.offset(0),
		transitionSet: copied,
		current:       1,
		finalStates:   fcopied,
	}
}
