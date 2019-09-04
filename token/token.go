package token

type Token interface {
	token()
}

type Char string

func (c Char) token() {}

func (c Char) char() {}

type Or string

func (o Or) token() {}

func (o Or) or() {}

type And string

func (a And) token() {}

func (a And) and() {}

type Asterisk string

func (a Asterisk) token() {}

func (a Asterisk) asterisk() {}

type LeftParenthesis string

func (l LeftParenthesis) token() {}

func (l LeftParenthesis) leftParenthesis() {}

type RightParenthesis string

func (r RightParenthesis) token() {}

func (r RightParenthesis) rightParenthesis() {}

type EOF string

func (e EOF) token() {}

func (e EOF) eof() {}
