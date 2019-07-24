package re

import "github.com/Salpadding/regexp/fsa"

type RegExp interface {
	fsa.FSA
	Match(string) bool
}

type regexp struct {
	fsa.FSA
}

func (r *regexp) Match(s string) bool {
	r.Reset()
	r.InputString(s)
	return r.IsAccept()
}

func Compile(re string) (RegExp, error) {
	nfa, err := fsa.New(re)
	if err != nil {
		return nil, err
	}
	return &regexp{
		nfa.ToDFA(),
	}, nil
}
