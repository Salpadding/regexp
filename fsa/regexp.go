package fsa

type RegExp interface {
	Input(...rune) RegExp
	InputString(string) RegExp
	IsAccept() bool
	Reset() RegExp
}
