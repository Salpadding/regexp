package fsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChar(t *testing.T) {
	n := NewChar('a')
	n.Input('a')
	assert.True(t, n.IsAccept())
}

func TestConcat(t *testing.T) {
	n1 := NewChar('a')
	n2 := NewChar('b')
	n3 := n1.concat(n2)
	n3.Input('a')
	assert.False(t, n3.IsAccept())
	n3.Input('b')
	assert.True(t, n3.IsAccept())
}

func TestConcatMore(t *testing.T) {
	n := NewChar('a').concat(NewChar('b')).concat(NewChar('c')).concat(NewChar('d'))
	n.Input('a')
	n.Input('b')
	n.Input('c')
	assert.False(t, n.IsAccept())
	n.Input('d')
	assert.True(t, n.IsAccept())
}

func TestOr(t *testing.T) {
	n1 := NewChar('a')
	n2 := NewChar('b')
	n3 := n1.or(n2)
	n3.Input('b')
	assert.True(t, n3.IsAccept())
}

func TestOrMore(t *testing.T) {
	n := NewChar('a').or(NewChar('b')).or(NewChar('c')).or(NewChar('d'))
	n.Input('d')
	assert.True(t, n.IsAccept())
}

func TestClosure(t *testing.T) {
	n := NewChar('a')
	n = n.kleen()
	n.InputString("aaaaaaa")
	assert.True(t, n.IsAccept())
	n = NewChar('a')
	n = n.kleen()
	n.InputString("")
	assert.True(t, n.IsAccept())
}

func TestOrClosure(t *testing.T) {
	n := NewChar('a').or(NewChar('b')).kleen()
	n.InputString("bbbaaabbbaaaaaaabbb")
	assert.True(t, n.IsAccept())
}

func TestConcatClosure(t *testing.T) {
	n := NewChar('a').concat(NewChar('b')).kleen()
	n.InputString("abababab")
	assert.True(t, n.IsAccept())
}

func TestClosureConcat(t *testing.T) {
	n := NewChar('b').kleen().concat(NewChar('a'))
	n.InputString("bbbbbbbbba")
	assert.True(t, n.IsAccept())

}

func TestOrClosureConcat(t *testing.T) {
	n := NewChar('a').or(NewChar('b')).kleen().concat(NewChar('c'))
	n.InputString("abababababaaaaabbbc")
	assert.True(t, n.IsAccept())
}
