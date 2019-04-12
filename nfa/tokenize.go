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
	or:               &token{code: token_or},
	closure:          &token{code: token_closure},
	leftParentheses:  &token{code: token_leftParentheses},
	rightParentheses: &token{code: token_rightParentheses},
}

var concat = &token{
	code: token_concat,
}

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
	return pretokenized
}
