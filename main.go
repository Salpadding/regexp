package main

import (
	"fmt"

	"github.com/Salpadding/regexp/re"
)

func main() {
	reg := re.Compile("( a | b ) * cd")
	reg.InputString("aaaaaabbbbbbbaaaaabbbcd")
	fmt.Println(reg.IsAccept())
}
