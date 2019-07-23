package fsa

import (
	"bufio"
	"io"
)

type code int

const (
	tokenChar    code = iota // character
	tokenConcat              // for build ast
	tokenOr                  // represent |
	tokenClosure             // kleen closure
	tokenLeftParentheses
	tokenRightParentheses
	tokenOneOrMore   // +
	tokenNoneOrOne   // ?
	tokenWildcard    // . match any ascii character
	tokenDigital     // \d match digital 0,1...9
	tokenLetters     // \w match letters a,b...z A,B...Z
	tokenNonDigital  // \D match non-digital character
	tokenNonLetter   // \W match non-letters
	tokenRange       // [a-z0-9] match character range
)

const (
	leftParentheses  = '('
	rightParentheses = ')'
	escape           = '\\'
	or               = '|'
	closure          = '*'
	whiteSpace       = ' '
	dot              = '.'
	plus             = '+'
	question         = '?'
)

type token struct {
	code       code
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
	plus:             {code: tokenOneOrMore},
	question:         {code: tokenNoneOrOne},
}

var escapes = map[rune]*token{
	's':  {code: tokenChar, value: ' '},
	't':  {code: tokenChar, value: '\t'},
	'n':  {code: tokenChar, value: '\n'},
	'\\': {code: tokenChar, value: '\\'},
	'w':  {code: tokenLetters},
	'd':  {code: tokenDigital},
	'W':  {code: tokenNonLetter},
	'D':  {code: tokenNonDigital},
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
		case whiteSpace, '\t':
			continue
		case escape:
			// TODO: \s \w \d \n ...
			r, _, err = runeReader.ReadRune()
			if err != nil {
				panic("unexpected eof")
			}
			esc, ok := escapes[r]
			if !ok {
				esc = &token{code: tokenChar, value: r}
			}
			pretokenized = append(pretokenized, esc)
		case leftParentheses, rightParentheses, or, closure, dot, plus, question:
			pretokenized = append(pretokenized, cache[r])
		case '[':
		default:
			pretokenized = append(pretokenized, &token{code: tokenChar, value: r})
		}
	}
	// insert concat token between
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
