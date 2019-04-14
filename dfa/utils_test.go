package dfa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegerSet(t *testing.T) {
	var finalStates integerSet = map[int]bool{
		2: true,
	}
	assert.False(t, finalStates.has(0))
	entries := finalStates.entries()
	assert.Equal(t, 2, entries[0])
	assert.Equal(t, 1, len(entries))
}
