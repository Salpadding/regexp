package dfa

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

func (s *tokenStack) push(tk *token) {
	s.data = append(s.data, tk)
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

func newStack(data []*token) *tokenStack {
	if data == nil {
		data = make([]*token, 0)
	}
	return &tokenStack{
		data: data,
		pc:   0,
	}
}

func traverse(tree *token, cb func(*token)) {
	if tree == nil {
		return
	}
	cb(tree)
	if tree.leftChild != nil {
		traverse(tree.leftChild, cb)
	}
	if tree.rightChild != nil {
		traverse(tree.rightChild, cb)
	}
}

func (tree *token) stack() *tokenStack {
	var tks []*token
	traverse(tree, func(tk *token) {
		tks = append([]*token{tk}, tks...)
	})
	return newStack(tks)
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
			return ntk
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
