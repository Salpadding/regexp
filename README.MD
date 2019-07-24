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
	re, err := re.Compile(`[a-z0-9A-Z]+@[a-z0-9A-Z]+\.[a-z0-9A-Z]+`)
	assert.NoError(t, err)
	re.InputString("m6567fc@outlook.com")
	assert.True(t, re.IsAccept())
	re.Reset()
	re.InputString("abbbbb@yyy")
	assert.False(t, re.IsAccept())
}

```