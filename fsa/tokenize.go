package fsa

import (
	"bufio"
	"io"
)

const (
	tokenChar = iota // character
	tokenConcat
	tokenOr      // represent |
	tokenClosure // kleen closure
	tokenLeftParentheses
	tokenRightParentheses
	tokenWildcard // . match any character
)

const (
	leftParentheses  = '('
	rightParentheses = ')'
	escape           = '\\'
	or               = '|'
	closure          = '*'
	whiteSpace       = ' '
	dot              = '.'
)

type token struct {
	code       int
	value      rune
	leftChild  *token
	rightChild *token
}

var cache = map[rune]*token{
	or:               {code: tokenOr, value: '|'},
	closure:          {code: tokenClosure, value: '*'},
	leftParentheses:  {code: tokenLeftParentheses, value: '('},
	rightParentheses: {code: tokenRightParentheses, value: ')'},
	dot:              {code: tokenWildcard},
}

var concat = &token{
	code:  tokenConcat,
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
			pretokenized = append(pretokenized, &token{code: tokenChar, value: r})
		case leftParentheses, rightParentheses, or, closure:
			pretokenized = append(pretokenized, cache[r])
		case dot:
			pretokenized = append(pretokenized, cache[r])
		default:
			pretokenized = append(pretokenized, &token{code: tokenChar, value: r})
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
		if token.code == tokenOr || next.code == tokenOr {
			continue
		}
		if next.code == tokenClosure {
			continue
		}
		if next.code == tokenRightParentheses {
			continue
		}
		if token.code == tokenLeftParentheses {
			continue
		}
		res = append(res, concat)
	}
	return res
}
