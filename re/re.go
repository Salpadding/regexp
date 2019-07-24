package re

import "github.com/Salpadding/regexp/fsa"

type RegExp fsa.FSA

func Compile(re string) (RegExp, error) {
	nfa, err := fsa.New(re)
	if err != nil {
		return nil, err
	}
	return nfa.ToDFA(), nil
}
