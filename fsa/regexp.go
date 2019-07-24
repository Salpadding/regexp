package fsa

type FSA interface {
	Input(...rune) FSA
	InputString(string) FSA
	IsAccept() bool
	Reset() FSA
}
