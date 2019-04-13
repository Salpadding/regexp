package nfa

import (
	"bufio"
	"io"
)

const (
	token_char = iota
	token_concat
	token_or
	token_closure
	token_leftParentheses
	token_rightParentheses
)

type token struct {
	code       int
	value      rune
	leftChild  *token
	rightChild *token
}

var cache = map[rune]*token{
	or:               &token{code: token_or, value: '|'},
	closure:          &token{code: token_closure, value: '*'},
	leftParentheses:  &token{code: token_leftParentheses, value: '('},
	rightParentheses: &token{code: token_rightParentheses, value: ')'},
}

var concat = &token{
	code:  token_concat,
	value: '+',
}

// TODO: keep parentheses closed always
func tokenize(reader io.Reader) []*token {
	var pretokenized []*token
	runeReader := bufio.NewReader(reader)
	for r, _, err := runeReader.ReadRune(); err == nil; r, _, err = runeReader.ReadRune() {
		switch r {
		case whiteSpace:
			continue
		case escape:
			// TODO: \s \w \d \n ...
			r, _, err = runeReader.ReadRune()
			if err != nil {
				panic("unexpected eof")
			}
			pretokenized = append(pretokenized, &token{code: token_char, value: r})
		case leftParentheses, rightParentheses, or, closure:
			pretokenized = append(pretokenized, cache[r])
		default:
			pretokenized = append(pretokenized, &token{code: token_char, value: r})
		}
	}
	// TODO: insert concat tokens
	var res []*token
	for i, token := range pretokenized {
		res = append(res, token)
		if i+1 == len(pretokenized) {
			break
		}
		next := pretokenized[i+1]
		if token.code == token_or || next.code == token_or {
			continue
		}
		if next.code == token_closure {
			continue
		}
		if next.code == token_rightParentheses {
			continue
		}
		if token.code == token_leftParentheses {
			continue
		}
		res = append(res, concat)
	}
	return res
}
