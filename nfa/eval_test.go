package nfa

import (
	"fmt"
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

// TODO: this is nfa not dfa
func TestEvalMore2(t *testing.T) {
	nfa := New("(a|b)*b")
	nfa.InputString("b")
	fmt.Println(nfa.currentStates.entries())
	assert.True(t, nfa.IsAccept())
}