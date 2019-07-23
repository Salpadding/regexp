package re

import "github.com/Salpadding/regexp/fsa"

type RegExp interface {
	Input(...rune)
	InputString(string)
	IsAccept() bool
}

func Compile(re string) RegExp {
	return fsa.New(re).ToDFA()
}
