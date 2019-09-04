package ast

type Expression interface {
	expression()
}

type Concat struct {
	Left  Expression
	Right Expression
}

func (c *Concat) expression() {}

type Or struct {
	Left  Expression
	Right Expression
}

func (o *Or) expression() {}

type Closure struct {
	Expression
}

type Char string

func (c Char) expression() {}

type Digital string

func (d Digital) expression() {}

type NonDigital string

func (n NonDigital) expression() {}

type WildCard string

func (w WildCard) expression() {}

type OneOrMore struct {
	Expression
}

type NoneOrOne struct {
	Expression
}

type Letters string

func (l Letters) expression() {}

type NonLetters string

func (n NonLetters) expression() {}

type Range map[rune]rune

func (r Range) expression() {}
