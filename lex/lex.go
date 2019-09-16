package lex

import (
	"errors"
	"io"

	"github.com/Salpadding/regexp/token"
)

var whiteSpaces = map[rune]bool{
	' ':  true,
	'\n': true,
	'\r': true,
	'\t': true,
}

var escapes = map[rune]token.Token{
	's':  token.Char(' '),
	't':  token.Char('\t'),
	'n':  token.Char('\n'),
	'\\': token.Char('\\'),
	'w':  token.Letters("[a-zA-Z]"),
	'd':  token.Digital("[0-9]"),
	'W':  token.NonDigital(`\W`),
	'D':  token.NonDigital(`\D`),
}

var operators = map[rune]token.Token{
	'|': token.Or('|'),
	'(': token.LeftParenthesis('('),
	')': token.RightParenthesis(')'),
	'*': token.Asterisk('*'),
	'+': token.Plus('+'),
	'?': token.QuestionMark('?'),
	'.': token.WildCard('.'),
	'[': token.LeftBracket('['),
	']': token.RightBracket(']'),
}

type Char interface {
	rune() rune
}

type char rune

func (c char) rune() rune {
	return rune(c)
}

func (c char) char() {}

type eof rune

func (e eof) rune() rune {
	return rune(e)
}

func (e eof) eof() {}

type Lexer struct {
	io.RuneReader
	current Char
	next    Char
}

func (l *Lexer) nextRune() Char {
	_, ok := l.current.(eof)
	if ok {
		return eof(0)
	}
	l.current = l.next
	r, _, err := l.RuneReader.ReadRune()
	if err != nil {
		l.next = eof(0)
		return l.current
	}
	l.next = char(r)
	return l.current
}

func New(reader io.RuneReader) *Lexer {
	l := &Lexer{
		RuneReader: reader,
	}
	l.nextRune()
	l.nextRune()
	return l
}

func (l *Lexer) NextToken() (token.Token, error) {
	// skip white spaces
	for {
		c, ok := l.current.(char)
		if !ok {
			break
		}
		if !whiteSpaces[rune(c)] {
			break
		}
		l.nextRune()
	}

	switch c := l.current.(type) {
	case eof:
		return token.EOF("EOF"), nil
	case char:
		r := rune(c)
		switch r {
		case '|', '*', '(', ')', '+', '?', '.':
			l.nextRune()
			return operators[r], nil
		case '\\':
			n, ok := l.next.(char)
			if !ok {
				return nil, errors.New("unexpected eof after slash")
			}
			tk, ok := escapes[rune(n)]
			if ok {
				l.nextRune()
				l.nextRune()
				return tk, nil
			}
			l.nextRune()
			l.nextRune()
			return token.Char(n), nil
		case '[':
			r := token.Ranges{}
			for {
				l.nextRune()
				_, ok := l.current.(eof)
				if ok {
					return nil, errors.New("unexpected eof")
				}
				if l.current.rune() == ']' {
					break
				}
				if l.next.rune() == '-' {
					n := l.current.rune()
					l.nextRune()
					l.nextRune()
					r[n] = l.current.rune()
					continue
				}
				r[l.current.rune()] = l.current.rune()
			}
			l.nextRune()
			return r, nil
		default:
			l.nextRune()
			return token.Char(r), nil
		}
	default:
		return nil, errors.New("invalid type")
	}
}
