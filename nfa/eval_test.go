package nfa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	nfa := New("( a | b ) * cd")
	nfa.InputString("aaaaaabbbbbbbaaaaabbbcd")
	assert.True(t, nfa.IsAccept())
}

func TestEvalMore(t *testing.T) {
	nfa := New("( a | b ) * cd")
	nfa.InputString("aaaaaabbbbbbbaaaaabbbcd")
	assert.True(t, nfa.IsAccept())
}
