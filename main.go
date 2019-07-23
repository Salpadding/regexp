package main

import (
	"fmt"
	nfa_v2 "github.com/Salpadding/regexp/nfa.v2"
)

func main() {
	nfa := nfa_v2.New("( a | b ) * cd")
	nfa.InputString("aaaaaabbbbbbbaaaaabbbcd")
	fmt.Println(nfa.IsAccept())
}
