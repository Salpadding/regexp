package main

import (
	"fmt"

	"github.com/Salpadding/regexp/dfa"
)

func main() {
	dfa := dfa.New("( a | b ) * cd")
	dfa.InputString("aaaaaabbbbbbbaaaaabbbcd")
	fmt.Println(dfa.IsAccept())
}
