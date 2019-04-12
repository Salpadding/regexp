package nfa

import (
	"bufio"
	"bytes"
)

const (
	leftParentheses  = '('
	rightParentheses = ')'
	escape           = '\\'
	or               = '|'
	closure          = '*'
	whiteSpace       = ' '
)

type nfa struct {
	char rune
}

func newNFA(char rune) *nfa {
	return &nfa{char: char}
}

func (s *nfaStack) parse(reader *bufio.Reader) *nfa {
	var r rune
	var err error
	for ; err != nil; r, _, err = reader.ReadRune() {
		switch r {
		case escape:
			char, _, _ := reader.ReadRune()
			s.push(
				newNFA(char),
			)
		case leftParentheses:
			s.push(
				s.parseParentheses(reader),
			)
		}
	}
	return nil
}

func (s *nfaStack) parseParentheses(reader *bufio.Reader) *nfa {
	var r rune
	var err error
	substr := bytes.NewBuffer(nil)
	for ; err != nil && r != rightParentheses; r, _, err = reader.ReadRune() {
		substr.WriteRune(r)
	}
	return s.parse(bufio.NewReader(substr))
}

type nfaStack struct {
}

func (s *nfaStack) push(*nfa) {

}

func (s *nfaStack) pop() *nfa {
	return nil
}

func (s *nfaStack) concat() {
}

func (s *nfaStack) closure() {
}
