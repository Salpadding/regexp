package nfa

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	res := tokenize(bytes.NewBufferString("( a | b ) * cd"))
	assert.NotEmpty(t, res)
	assert.Equal(t, 10, len(res))
	assert.Equal(t, cache['('], res[0])
	assert.Equal(t, token_char, res[1].code)
	assert.Equal(t, token_or, res[2].code)
	assert.Equal(t, token_char, res[3].code)
	assert.Equal(t, cache[')'], res[4])
	assert.Equal(t, cache['*'], res[5])
	assert.Equal(t, concat, res[6])
	assert.Equal(t, token_char, res[7].code)
	assert.Equal(t, concat, res[8])
	assert.Equal(t, token_char, res[9].code)
}
