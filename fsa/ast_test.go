package fsa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildAST(t *testing.T) {
	res, err := tokenize("( a | b ) * cd")
	assert.NoError(t, err)
	tree := parse(res)
	assert.NotNil(t, tree)
}

func TestBuildAST2(t *testing.T) {
	res, err := tokenize("()")
	assert.NoError(t, err)
	tree := parse(res)
	assert.NotNil(t, tree)
}

// TestBuildAST3 covers boundary conditions
func TestBuildAST3(t *testing.T) {
	res, err := tokenize("(a|b|c|d) | e ( h* (i| j) k)")
	assert.NoError(t, err)
	tree := parse(res)
	assert.NotNil(t, tree)
}

func TestBuildAST4(t *testing.T) {
	res, err := tokenize("a*")
	assert.NoError(t, err)
	tree := parse(res)
	assert.NotNil(t, tree)
}

func TestBuildAST5(t *testing.T) {
	res, err := tokenize("ab|cd")
	assert.NoError(t, err)
	parse(res)
	fmt.Println("===========")
}

func TestBuildAST6(t *testing.T) {
	res, err := tokenize(`(".*")|('.*')`)
	assert.NoError(t, err)
	tree := parse(res)
	assert.NotNil(t, tree)
	fmt.Println("===========")
}
