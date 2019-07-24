package main

import (
	"fmt"

	"github.com/Salpadding/regexp/re"
)

func main() {
	reg, _ := re.Compile("(a|b)*cd")
	reg.InputString("aaaaaabbbbbbbaaaaabbbcd")
	fmt.Println(reg.IsAccept()) // true, fulfilled
	reg.InputString("----")
	fmt.Println(reg.IsAccept()) // false

	// match golang keywords
	keywordMatcher, _ := re.Compile(
		`go|interface|struct|func|import` +
			`|package|type|const|if|range`,
	)
	keywordMatcher.InputString("interface")
	fmt.Println(keywordMatcher.IsAccept()) // true

	r, _ := re.Compile("ab|cd")
	r.InputString("cd")
	fmt.Println(r.IsAccept()) // true

	// match string
	stringMatcher, _ := re.Compile(`(".*")|('.*')`)
	stringMatcher.InputString(`"this is a string ;;'' dsfaf"`)
	fmt.Println(stringMatcher.IsAccept()) // true

	stringMatcher.Reset()
	stringMatcher.InputString(`'this is a string ;;"" dsfaf'`)
	fmt.Println(stringMatcher.IsAccept()) // true

	stringMatcher.Reset()
	stringMatcher.InputString(`'this is a string ;;"" dsfaf`)
	fmt.Println(stringMatcher.IsAccept()) // false
}
