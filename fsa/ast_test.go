package fsa

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildAST(t *testing.T) {
	res := tokenize(bytes.NewBufferString("( a | b ) * cd"))
	tree := buildAST(&tokenStack{
		data: res,
		pc:   0,
	}, nil)
	assert.NotNil(t, tree)
	assert.Equal(t, tokenConcat, tree.code)
	assert.Equal(t, tokenClosure, tree.leftChild.code)
	assert.Equal(t, tokenOr, tree.leftChild.leftChild.code)
	assert.Nil(t, tree.leftChild.rightChild)
	assert.Equal(t, tokenConcat, tree.rightChild.code)
	assert.Equal(t, 'c', tree.rightChild.leftChild.value)
}

func TestBuildAST2(t *testing.T) {
	res := tokenize(bytes.NewBufferString("()"))
	tree := buildAST(&tokenStack{
		data: res,
		pc:   0,
	}, nil)
	assert.Nil(t, tree)
}

// TestBuildAST3 covers boundary conditions
func TestBuildAST3(t *testing.T) {
	res := tokenize(bytes.NewBufferString("(a|b|c|d) | e ( h* (i| j) k)"))
	tree := buildAST(&tokenStack{
		data: res,
		pc:   0,
	}, nil)
	assert.NotNil(t, tree)
	assert.Equal(t, tokenOr, tree.code)
	assert.Equal(t, tokenClosure, tree.rightChild.rightChild.leftChild.code)
	assert.Equal(t, 'k', tree.rightChild.rightChild.rightChild.rightChild.value)
	assert.Equal(t, 'a', tree.leftChild.leftChild.value)
}

func TestBuildAST4(t *testing.T) {
	res := tokenize(bytes.NewBufferString("a*"))
	tree := buildAST(&tokenStack{
		data: res,
		pc:   0,
	}, nil)
	assert.NotNil(t, tree)
	assert.Equal(t, tokenClosure, tree.code)
}
