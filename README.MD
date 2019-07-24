# Regexp

a regular expression engine based on non-deterministic finite automata.

only supports ascii character 

```go
package main

import (
	"fmt"

	"github.com/Salpadding/regexp/re"
)

func main() {
	reg := re.Compile("( a | b ) * cd")
	reg.InputString("aaaaaabbbbbbbaaaaabbbcd")
	fmt.Println(reg.IsAccept())
	reg.InputString("----")
	fmt.Println(reg.IsAccept())
}

```