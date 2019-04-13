package nfa

import (
	"errors"
)

type tokenStack struct {
	data []*token
	pc   int
}

func (s *tokenStack) pop() (*token, error) {
	t, err := s.peek()
	if err != nil {
		return nil, err
	}
	s.pc++
	return t, nil
}

func (s *tokenStack) peek() (*token, error) {
	if s.pc >= len(s.data) {
		return nil, errors.New("eof")
	}
	t := s.data[s.pc]
	return t, nil
}

func (s *tokenStack) shift(idx int) {
	s.pc += idx
}

func buildAST(s *tokenStack, left *token) *token {
	l := left
	var r *token
	for tk, err := s.pop(); err == nil; tk, err = s.pop() {
		switch tk.code {
		case token_leftParentheses:
			l = buildAST(s, nil)
		case token_rightParentheses:
			return l
		case token_closure:
			tk, err = s.pop()
			ntk := &token{
				code:      token_closure,
				leftChild: l,
			}
			if err == nil && tk.code == token_concat {
				s.shift(-1)
				return buildAST(s, ntk)
			}
			return l
		case token_concat:
			r = buildAST(s, nil)
			if r == nil {
				return l
			}
			return &token{
				code:       token_concat,
				leftChild:  l,
				rightChild: r,
			}
		case token_or:
			r = buildAST(s, nil)
			if r == nil {
				return l
			}
			return &token{
				code:       token_or,
				leftChild:  l,
				rightChild: r,
			}
		default:
			l = tk
		}
	}
	return l
}
