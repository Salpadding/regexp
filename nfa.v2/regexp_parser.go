package nfa_v2

import (
	"bytes"
)

type nfaStack []*NFA

func (s *nfaStack) push(n *NFA) {
	*s = append(*s, n)
}

func (s *nfaStack) pop() *NFA {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func New(regexp string) *NFA {
	res := tokenize(bytes.NewBufferString(regexp))
	tree := buildAST(&tokenStack{
		data: res,
		pc:   0,
	}, nil)
	stack := tree.stack()
	var nfaStack nfaStack = make([]*NFA, 0)
	for tk, err := stack.pop(); err == nil; tk, err = stack.pop() {
		switch tk.code {
		case tokenClosure:
			nfaStack.push(
				nfaStack.pop().kleen(),
			)
		case tokenConcat:
			a := nfaStack.pop()
			b := nfaStack.pop()
			nfaStack.push(a.concat(b))
		case tokenOr:
			a := nfaStack.pop()
			b := nfaStack.pop()
			nfaStack.push(a.or(b))
		case tokenChar:
			nfaStack.push(
				NewChar(tk.value),
			)
		default:
			panic("unexpected token")
		}
	}
	return nfaStack.pop()
}
