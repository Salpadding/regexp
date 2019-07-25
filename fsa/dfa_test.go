package fsa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDFANewChar(t *testing.T) {
	n := NewChar('a').ToDFA()
	n.Input('a')
	assert.True(t, n.IsAccept())
}

func TestDFAConcat(t *testing.T) {
	n1 := NewChar('a')
	n2 := NewChar('b')
	n3 := n1.concat(n2).ToDFA()
	n3.Input('a')
	assert.False(t, n3.IsAccept())
	n3.Input('b')
	assert.True(t, n3.IsAccept())
}

func TestDFAConcatMore(t *testing.T) {
	n := NewChar('a').concat(NewChar('b')).concat(NewChar('c')).concat(NewChar('d')).ToDFA()
	n.Input('a')
	n.Input('b')
	n.Input('c')
	assert.False(t, n.IsAccept())
	n.Input('d')
	assert.True(t, n.IsAccept())
}

func TestDFAOr(t *testing.T) {
	n1 := NewChar('a')
	n2 := NewChar('b')
	n3 := n1.or(n2).ToDFA()
	n3.Input('b')
	assert.True(t, n3.IsAccept())
}

func TestDFAOrMore(t *testing.T) {
	n := NewChar('a').or(NewChar('b')).or(NewChar('c')).or(NewChar('d')).ToDFA()
	n.Input('d')
	assert.True(t, n.IsAccept())
	n.Input('a')
	assert.False(t, n.IsAccept())
}

func TestDFAClosure(t *testing.T) {
	n := NewChar('a').kleen().ToDFA()
	n.InputString("aaaaaaa")
	assert.True(t, n.IsAccept())
	n = NewChar('a').kleen().ToDFA()
	n.InputString("")
	assert.True(t, n.IsAccept())
}

func TestDFAOrClosure(t *testing.T) {
	n := NewChar('a').or(NewChar('b')).kleen().ToDFA()
	n.InputString("bbbaaabbbaaaaaaabbb")
	assert.True(t, n.IsAccept())
}

func TestDFAConcatClosure(t *testing.T) {
	n := NewChar('a').concat(NewChar('b')).kleen().ToDFA()
	n.InputString("abababab")
	assert.True(t, n.IsAccept())
}

func TestDFAClosureConcat(t *testing.T) {
	n := NewChar('b').kleen().concat(NewChar('a')).ToDFA()
	n.InputString("bbbbbbbbba")
	assert.True(t, n.IsAccept())

}

func TestDFAOrClosureConcat(t *testing.T) {
	n := NewChar('a').or(NewChar('b')).kleen().concat(NewChar('c')).ToDFA()
	n.InputString("abababababaaaaabbbc")
	assert.True(t, n.IsAccept())
}

func TestWildcard(t *testing.T) {
	n := newWildCard().kleen().ToDFA()
	n.InputString("abcdafafjofaj-2r02-f]2vgjadv;gkdfvamjadofff  dsfgkkk  -")
	assert.True(t, n.IsAccept())
	n = newWildCard().concat(NewChar('a')).ToDFA()
	n.InputString("ca")
	assert.True(t, n.IsAccept())
	n = newWildCard().concat(NewChar('a')).ToDFA()
	n.InputString("cb")
	assert.False(t, n.IsAccept())
}

func TestRefine(t *testing.T) {
	d := &DFA{
		transitions: map[rune]map[state]state{
			'a': {0: 1, 1: 1, 4: 0, 2: 1, 3: 3, 5: 5,},
			'b': {1: 4, 0: 2, 2: 3, 3: 2, 4: 5, 5: 4,},
		},
		finalStates:  newStateSet(0, 1),
		maximumState: 5,
	}

	d = d.Minimize()
	fmt.Println("====")
}

func TestMinimize1(t *testing.T) {
	nfa, err := New(`wisdom://([0-9a-f]+@)?((\d+\.\d+\.\d+\.\d+)|[0-9a-zA-Z]+)(:[0-9]+)?`)
	assert.NoError(t, err)
	dfa1 := nfa.ToDFA()
	dfa2 := dfa1.Minimize()
	assert.True(t, dfa2.maximumState < dfa1.maximumState)
}

func TestMinimize2(t *testing.T) {
	nfa, err := New(`-?[0-9]+(\.[0-9]+)?`)
	assert.NoError(t, err)
	dfa1 := nfa.ToDFA()
	dfa2 := dfa1.Minimize()
	assert.True(t, dfa2.maximumState < dfa1.maximumState)
}
