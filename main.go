package main

import (
	"fmt"

	"github.com/Salpadding/regexp/nfa"
)

func main() {
	nfa := nfa.New("( a | b ) * cd")
	nfa.InputString("aaaaaabbbbbbbaaaaabbbcd")
	fmt.Println(nfa.IsAccept())
}
