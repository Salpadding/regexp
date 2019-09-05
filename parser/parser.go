package parser

import (
	"errors"
	"io"

	"github.com/Salpadding/regexp/ast"
	"github.com/Salpadding/regexp/lex"
	"github.com/Salpadding/regexp/token"
)

type Precedence int

const (
	_ Precedence = iota
	lowest
	or
	and
	postfix
)

func (p *Parser) needConcat(t token.Token) bool {
	switch t.(type) {
	case token.NonDigital, token.Digital, token.Letters,
		token.NonLetters, token.RightParenthesis, token.QuestionMark,
		token.Plus, token.Asterisk, token.Char, token.Ranges:
		return true
	default:
		return false
	}
}

func (p *Parser) isInfix(t token.Token) bool {
	switch t.(type) {
	case token.Or, token.RightParenthesis, token.EOF, token.Plus, token.Asterisk, token.QuestionMark:
		return true
	default:
		return false
	}
}

type Parser struct {
	*lex.Lexer
	current token.Token
	next    token.Token
	cache   []token.Token
}

func (p *Parser) nextToken() (token.Token, error) {
	var err error
	p.current = p.next
	if len(p.cache) > 0 {
		p.next = p.cache[0]
		p.cache = p.cache[:len(p.cache)-1]
		return p.current, nil
	}
	next, err := p.Lexer.NextToken()
	if err != nil {
		return nil, err
	}

	if p.needConcat(p.current) && !p.isInfix(next) {
		p.next = token.And("")
		p.cache = append(p.cache, next)
		return p.current, nil
	}
	p.next = next
	return p.current, nil
}

func New(reader io.RuneReader) (*Parser, error) {
	p := &Parser{
		Lexer: lex.New(reader),
		cache: []token.Token{},
	}
	if _, err := p.nextToken(); err != nil {
		return nil, err
	}
	if _, err := p.nextToken(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Parser) precedence(tk token.Token) Precedence {
	switch tk.(type) {
	case token.And:
		return and
	case token.Asterisk, token.Plus, token.QuestionMark:
		return postfix
	case token.Or:
		return or
	default:
		return 0
	}
}

func (p *Parser) Parse() (ast.Expression, error) {
	return p.parsePrecedence(lowest)
}

func (p *Parser) parsePrecedence(precedence Precedence) (ast.Expression, error) {
	left, err := p.parsePrefix()
	if err != nil {
		return nil, err
	}
	for {
		pd := p.precedence(p.current)
		if precedence >= pd {
			break
		}
		left, err = p.parseInfix(left, p.current)
		if err != nil {
			return nil, err
		}
	}
	return left, nil
}

func (p *Parser) parseInfix(left ast.Expression, op token.Token) (ast.Expression, error) {
	switch op.(type) {
	case token.And:
		if _, err := p.nextToken(); err != nil {
			return nil, err
		}
		right, err := p.parsePrecedence(and)
		if err != nil {
			return nil, err
		}
		return &ast.Concat{
			Left:  left,
			Right: right,
		}, nil
	case token.Asterisk:
		if _, err := p.nextToken(); err != nil {
			return nil, err
		}
		return &ast.Closure{Expression: left}, nil
	case token.Plus:
		if _, err := p.nextToken(); err != nil {
			return nil, err
		}
		return &ast.OneOrMore{Expression: left}, nil
	case token.QuestionMark:
		if _, err := p.nextToken(); err != nil {
			return nil, err
		}
		return &ast.NoneOrOne{Expression: left}, nil
	case token.Or:
		if _, err := p.nextToken(); err != nil {
			return nil, err
		}
		right, err := p.parsePrecedence(or)
		if err != nil {
			return nil, err
		}
		return &ast.Or{
			Left:  left,
			Right: right,
		}, nil
	default:
		return nil, errors.New("unexpected operator")
	}
}

func (p *Parser) parsePrefix() (ast.Expression, error) {
	current := p.current
	_, err := p.nextToken()
	if err != nil {
		return nil, err
	}
	switch c := current.(type) {
	case token.Char:
		return ast.Char(c), nil
	case token.NonLetters:
		return ast.NonLetters(c), nil
	case token.Letters:
		return ast.Letters(c), nil
	case token.Digital:
		return ast.Digital(c), nil
	case token.NonDigital:
		return ast.NonDigital(c), nil
	case token.LeftParenthesis:
		exp, err := p.parsePrecedence(lowest)
		if err != nil {
			return nil, err
		}
		_, err = p.nextToken()
		if err != nil {
			return nil, err
		}
		return exp, nil
	case token.Ranges:
		return ast.Ranges(c), nil
	default:
		return nil, errors.New("invalid token found at beginning")
	}
}
