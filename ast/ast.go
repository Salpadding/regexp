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
