package fsa

import (
	"bytes"
	"errors"
	"fmt"
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
	tab              = '\t'
)

type token struct {
	code   code
	value  rune
	ranges map[rune]rune //
}

var cache = map[rune]*token{
	or:               {code: tokenOr},
	closure:          {code: tokenClosure},
	leftParentheses:  {code: tokenLeftParentheses},
	rightParentheses: {code: tokenRightParentheses},
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
	code: tokenConcat,
}

// the range must match ((.-.)|.)+
var rangeMatcher = newWildCard().concat(
	NewChar('-')).concat(newWildCard(),
).or(newWildCard()).oneOrMore().ToDFA()

// TODO: keep parentheses closed always
func tokenize(program string) ([]*token, error) {
	var res []*token
	runes := []rune(program)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		switch r {
		case whiteSpace, tab:
			continue
		case escape:
			if i+1 >= len(runes) {
				return nil, fmt.Errorf("unexpected eof after %s", string(runes[:i]))
			}
			esc, ok := escapes[runes[i+1]]
			if !ok {
				esc = &token{code: tokenChar, value: runes[i+1]}
			}
			res = append(res, esc)
			i++
		case leftParentheses, rightParentheses, or, closure, dot, plus, question:
			res = append(res, cache[r])
		case '[':
			buf := bytes.NewBuffer(nil)
			i++
			for {
				if i >= len(runes) {
					return nil, fmt.Errorf("unexpected eof after %s", string(runes[:i]))
				}
				if runes[i] == ']' {
					break
				}
				buf.WriteRune(runes[i])
				i++
			}
			t, err := readRange(buf.String())
			if err != nil {
				return nil, err
			}
			res = append(res, t)
		default:
			res = append(res, &token{code: tokenChar, value: r})
		}
	}
	return res, nil
}

func readRange(buf string) (*token, error) {
	rangeMatcher.Reset()
	rangeMatcher.InputString(buf)
	ranges := map[rune]rune{}
	if !rangeMatcher.IsAccept() {
		return nil, errors.New("invalid range " + buf)
	}
	runes := []rune(buf)
	for i := 0; i < len(runes); {
		if i+1 == len(runes) {
			ranges[runes[i]] = epsilon
			break
		}
		if runes[i+1] == '-' {
			ranges[runes[i]] = runes[i+2]
			i += 3
			continue
		}
		ranges[runes[i]] = epsilon
		i++
	}
	return &token{
		code:   tokenRange,
		ranges: ranges,
	}, nil
}
