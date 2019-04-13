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
	states      map[int]int           // states represented by graph
	transitions map[int][]*transition // transation functions of this automata
	current     int                   // current state of the nfa, typicallly initialzed 1
}

type transition struct {
	from int
	to   int
	char rune
}

func newNFA(char rune) *nfa {
	return nil
}

func (n *nfa) concat(n2 *nfa) *nfa {
	return nil
}
