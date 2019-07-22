package nfa_v2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRune(t *testing.T) {
	n := NewRune('a')
	n.Input('a')
	assert.True(t, n.Accepted())
	n = NewRune('a')
	n.Input('b')
	assert.True(t, n.rejected)
	assert.False(t, n.Accepted())
}
