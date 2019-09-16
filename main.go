package main

import (
	"fmt"

	"github.com/Salpadding/regexp/re"
)

func main() {
	r, _ := re.Compile(`[a-z0-9A-Z]+@[a-z0-9A-Z]+\.[a-z0-9A-Z]+`)
	fmt.Println(r.Match("m6567fc@outlook.com"))
	fmt.Println(r.Match("abbbbb@yyy"))
}
