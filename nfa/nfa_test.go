package nfa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNFA(t *testing.T) {
	n := newNFA('a')
	n.Input('a')
	assert.True(t, n.IsAccept())
}

func TestConcat(t *testing.T) {
	n1 := newNFA('a')
	n2 := newNFA('b')
	n3 := n1.concat(n2)
	n3.Input('a')
	assert.False(t, n3.IsAccept())
	n3.Input('b')
	assert.True(t, n3.IsAccept())
}

func TestConcatMore(t *testing.T) {
	n := newNFA('a').concat(newNFA('b')).concat(newNFA('c')).concat(newNFA('d'))
	n.Input('a')
	n.Input('b')
	n.Input('c')
	assert.False(t, n.IsAccept())
	n.Input('d')
	assert.True(t, n.IsAccept())
}

func TestOr(t *testing.T) {
	n1 := newNFA('a')
	n2 := newNFA('b')
	n3 := n1.or(n2)
	n3.Input('b')
	assert.True(t, n3.IsAccept())
}

func TestOrMore(t *testing.T) {
	n := newNFA('a').or(newNFA('b')).or(newNFA('c')).or(newNFA('d'))
	n.Input('d')
	assert.True(t, n.IsAccept())
}

func TestClosure(t *testing.T) {
	n := newNFA('a')
	n = n.closure()
	n.InputString("aaaaaaa")
	assert.True(t, n.IsAccept())
	n = newNFA('a')
	n = n.closure()
	n.InputString("")
	assert.True(t, n.IsAccept())
}

func TestOrClosure(t *testing.T) {
	n := newNFA('a').or(newNFA('b')).closure()
	n.InputString("bbbaaabbbaaaaaaabbb")
	assert.True(t, n.IsAccept())
}

func TestConcatClosure(t *testing.T) {
	n := newNFA('a').concat(newNFA('b')).closure()
	n.InputString("abababab")
	assert.True(t, n.IsAccept())
}

func TestClosureConcat(t *testing.T) {
	n := newNFA('b').closure().concat(newNFA('a'))
	n.InputString("bbbbbbbbba")
	assert.True(t, n.IsAccept())

}

func TestOrClosureConcat(t *testing.T) {
	n := newNFA('a').or(newNFA('b')).closure().concat(newNFA('c'))
	assert.True(t, n.finalStates.has(4))
	assert.Equal(t, 1, len(n.finalStates.entries()))
}
