package nfa

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	res := tokenize(bytes.NewBufferString("( a | b ) * cd"))
	assert.NotEmpty(t, res)
	assert.Equal(t, 8, len(res))
	assert.Equal(t, cache['('], res[0])
}
