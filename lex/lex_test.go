package lex

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/Salpadding/regexp/token"
)

func Test(t *testing.T) {
	l := New(bytes.NewBufferString("(a|b|c|d) | e ( h* (i| j) k)"))
	var tks []token.Token
	for{
		tk, err := l.NextToken()
		if err != nil{
			t.Error(err)
		}
		if _, ok := tk.(token.EOF); ok{
			break
		}
		tks = append(tks, tk)
	}
	fmt.Printf("%v", tks)
}
