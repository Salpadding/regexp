package main

import (
	"fmt"

	"github.com/Salpadding/regexp/re"
)

func main() {
	reg := re.Compile("( a | b ) * cd")
	reg.InputString("aaaaaabbbbbbbaaaaabbbcd")
	fmt.Println(reg.IsAccept()) // true, fulfilled
	reg.InputString("----")
	fmt.Println(reg.IsAccept()) // false

	// match keyword
	keywordMatcher := re.Compile("(go)|(interface)|(struct)|(func)|(import)|(package)|(type)|(const)")
	keywordMatcher.InputString("interface")
	fmt.Println(keywordMatcher.IsAccept()) // true

	r := re.Compile("(ab)|(cd)")
	r.InputString("cd")
	fmt.Println(r.IsAccept()) // true
}
