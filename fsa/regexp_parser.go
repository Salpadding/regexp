package fsa

import (
	"bytes"

	"github.com/Salpadding/regexp/ast"
	"github.com/Salpadding/regexp/parser"
)

func New(regexp string) (*NFA, error) {
	ps, err := parser.New(bytes.NewBufferString(regexp))
	if err != nil {
		return nil, err
	}
	exp, err := ps.Parse()
	if err != nil {
		return nil, err
	}
	return Eval(exp), nil
}

func Eval(expr ast.Expression) *NFA {
	switch e := expr.(type) {
	case *ast.Concat:
		return Eval(e.Left).concat(Eval(e.Right))
	case *ast.Or:
		return Eval(e.Left).or(Eval(e.Right))
	case *ast.Closure:
		return Eval(e.Expression).kleen()
	case *ast.NoneOrOne:
		return Eval(e.Expression).noneOrOne()
	case *ast.OneOrMore:
		return Eval(e.Expression).oneOrMore()
	case ast.Char:
		return NewChar(rune(e[0]))
	case ast.Digital:
		return newDigital()
	case ast.Letters:
		return newLetters()
	case ast.WildCard:
		return newWildCard()
	case ast.NonDigital:
		return newNonDigital()
	case ast.NonLetters:
		return newNonLetter()
	case ast.Ranges:
		n := newEmpty()
		for k, v := range e {
			if v == epsilon {
				n.addTransition(k, 0, 1)
				continue
			}
			for r := k; r <= v; r++ {
				n.addTransition(r, 0, 1)
			}
		}
		return n
	}
	return nil
}
