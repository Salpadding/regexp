package nfa

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildAST(t *testing.T) {
	res := tokenize(bytes.NewBufferString("( a | b )b"))
	tree := buildAST(&tokenStack{
		data: res,
		pc:   0,
	})
	assert.NotNil(t, tree)
	assert.Equal(t, token_concat, tree.code)
}
