package nfa

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

type nfa struct {
	states        *states        // states represented by graph
	transitionSet *transitionSet // transation functions of this automata
	currentStates *integerSet    // current state of the nfa, typicallly initialzed 1
	finalStates   *integerSet
	initialState  int
}

type transition interface {
	setTo(int)
	to() int
	guard(rune) bool
	copy() transition
}

type charTransition struct {
	t    int
	char rune
}

func (c *charTransition) to() int {
	return c.t
}

func (c *charTransition) setTo(in int) {
	c.t = in
}

func (c *charTransition) guard(in rune) bool {
	return c.char == in
}

func (c *charTransition) copy() transition {
	return &charTransition{
		t:    c.t,
		char: c.char,
	}
}

// [1] -- a --> [[2]]
func newNFA(char rune) *nfa {
	var finalStates integerSet = map[int]bool{
		2: true,
	}
	var currentStates integerSet = map[int]bool{
		1: true,
	}
	var ts transitions = []transition{
		&charTransition{
			t:    2,
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
		currentStates: &currentStates,
		finalStates:   &finalStates,
		initialState:  1,
	}
}

func (n *nfa) Input(r rune) {
	if n.currentStates.has(-1) && n.currentStates.len() == 1 {
		return
	}
	var newStates integerSet = make(map[int]bool)
	flag := false
	for _, state := range n.currentStates.entries() {
		for _, t := range *n.transitionSet.get(state) {
			if t.guard(r) {
				newStates.add(t.to())
				flag = true
			}
		}
	}
	if !flag {
		newStates.add(-1)
	}
	n.currentStates = &newStates
}

func (n *nfa) Reset() {
	var newStates integerSet = make(map[int]bool)
	n.currentStates = &newStates
}

func (n *nfa) InputString(s string) {
	for _, r := range s {
		n.Input(r)
	}
}

func (n *nfa) IsAccept() bool {
	for _, state := range n.currentStates.entries() {
		if n.finalStates.has(state) {
			return true
		}
	}
	return false
}

// [1] -- a --> [[2]] + [1] -- b --> [[2]] ===> [1] -- a --> [2] -- b --> [[3]]
func (n *nfa) concat(n2 *nfa) *nfa {
	// 核心思想是增加把 1 的 final states 作为 2 的初始状态
	var currentStates integerSet = map[int]bool{
		1: true,
	}
	// append offsets
	offset := n.states.len() - 1
	copied2 := n2.transitionSet.offset(0)
	t1 := copied2.remove(1)
	t1 = t1.offset(offset)
	copied1 := n.transitionSet.offset(0)

	for _, k := range n.finalStates.entries() {
		copied1.add(k, t1)
	}
	return &nfa{
		states: n.states.concat(
			n2.states.slice(1, n2.states.len()).offset(offset),
		),
		transitionSet: copied1.union(copied2.offset(offset)),
		currentStates: &currentStates,
		finalStates:   n2.finalStates.offset(offset),
	}
}

func (n *nfa) or(n2 *nfa) *nfa {

	var currentStates integerSet = map[int]bool{
		1: true,
	}

	// 核心思想是取两个final state 为并集, 把 2 的初始状态对应的 trans 加上 offset 给1
	offset := n.states.len() - 1

	copied := n2.transitionSet.copy()
	t1 := copied.remove(1)
	t1 = t1.offset(offset)
	copied = copied.offset(offset).add(1, t1)
	newFinals := n.finalStates.union(
		n2.finalStates.offset(offset),
	)
	return &nfa{
		states: n.states.concat(
			n2.states.slice(1, n2.states.len()).offset(offset),
		),
		transitionSet: n.transitionSet.union(
			copied,
		),
		currentStates: &currentStates,
		finalStates:   newFinals,
	}
}

func (n *nfa) closure() *nfa {
	var currentStates integerSet = map[int]bool{
		1: true,
	}
	copied := n.transitionSet.offset(0)
	for _, k := range n.finalStates.entries() {
		copied = copied.add(k, n.transitionSet.get(1))
	}
	fcopied := n.finalStates.offset(0)
	fcopied.add(1)
	return &nfa{
		states:        n.states.offset(0),
		transitionSet: copied,
		currentStates: &currentStates,
		finalStates:   fcopied,
	}
}
