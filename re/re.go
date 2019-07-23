package re

import "github.com/Salpadding/regexp/fsa"

func Compile(re string) fsa.RegExp {
	return fsa.New(re).ToDFA()
}
