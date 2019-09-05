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

type Plus string

func (p Plus) token() {}

func (p Plus) plus() {}

type QuestionMark string

func (q QuestionMark) token() {}

func (q QuestionMark) questionMark() {}

type WildCard string

func (w WildCard) token() {}

func (w WildCard) wildCard() {}

type Letters string

func (l Letters) token() {}

func (l Letters) letters() {}

type Digital string

func (d Digital) token() {}

func (d Digital) digital() {}

type NonLetters string

func (l NonLetters) token() {}

func (l NonLetters) nonLetters() {}

type NonDigital string

func (d NonDigital) token() {}

func (d NonDigital) nonDigital() {}

type LeftBracket string

func (l LeftBracket) token() {}

func (l LeftBracket) leftBracket() {}

type RightBracket string

func (r RightBracket) token() {}

func (r RightBracket) rightBracket() {}

type Ranges map[rune]rune

func (r Ranges) token() {}

func (r Ranges) ranges() {}
