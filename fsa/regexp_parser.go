package fsa

import (
	"bytes"
	"fmt"
)

func New(regexp string) *NFA {
	res := tokenize(bytes.NewBufferString(regexp))
	tree := buildAST(&tokenStack{
		data: res,
		pc:   0,
	}, nil)
	return eval(tree)
}

func eval(tree *token) *NFA {
	switch tree.code {
	case tokenChar:
		return NewChar(tree.value)
	case tokenConcat:
		return eval(tree.leftChild).concat(eval(tree.rightChild))
	case tokenOr:
		return eval(tree.leftChild).or(eval(tree.rightChild))
	case tokenClosure:
		return eval(tree.leftChild).kleen()
	default:
		panic(fmt.Sprintf("unexpected token type %d", tree.code))
	}
}
