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
