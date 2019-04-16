package nfa

import (
	"bytes"
)

type nfaStack []*nfa

func (s *nfaStack) push(n *nfa) {
	(*s) = append(*s, n)
}

func (s *nfaStack) pop() *nfa {
	top := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return top
}

func New(regexp string) *nfa {
	res := tokenize(bytes.NewBufferString(regexp))
	tree := buildAST(&tokenStack{
		data: res,
		pc:   0,
	}, nil)
	stack := tree.stack()
	var nfaStack nfaStack = make([]*nfa, 0)
	for tk, err := stack.pop(); err == nil; tk, err = stack.pop() {
		switch tk.code {
		case token_closure:
			nfaStack.push(
				nfaStack.pop().closure(),
			)
		case token_concat:
			a := nfaStack.pop()
			b := nfaStack.pop()
			nfaStack.push(a.concat(b))
		case token_or:
			a := nfaStack.pop()
			b := nfaStack.pop()
			nfaStack.push(a.or(b))
		case token_char:
			nfaStack.push(
				newNFA(tk.value),
			)
		default:
			panic("unexpected token")
		}
	}
	return nfaStack.pop()
}
