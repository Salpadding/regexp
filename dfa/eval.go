package dfa

import (
	"bytes"
)

type dfaStack []*dfa

func (s *dfaStack) push(n *dfa) {
	(*s) = append(*s, n)
}

func (s *dfaStack) pop() *dfa {
	top := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return top
}

func New(regexp string) *dfa {
	res := tokenize(bytes.NewBufferString(regexp))
	tree := buildAST(&tokenStack{
		data: res,
		pc:   0,
	}, nil)
	stack := tree.stack()
	var dfaStack dfaStack = make([]*dfa, 0)
	for tk, err := stack.pop(); err == nil; tk, err = stack.pop() {
		switch tk.code {
		case token_closure:
			dfaStack.push(
				dfaStack.pop().closure(),
			)
		case token_concat:
			a := dfaStack.pop()
			b := dfaStack.pop()
			dfaStack.push(a.concat(b))
		case token_or:
			a := dfaStack.pop()
			b := dfaStack.pop()
			dfaStack.push(a.or(b))
		case token_char:
			dfaStack.push(
				newNFA(tk.value),
			)
		default:
			panic("unexpected token")
		}
	}
	return dfaStack.pop()
}
