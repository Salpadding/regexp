package parser

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T){
	p, err := New(bytes.NewBufferString("(a|b|c|d) | e ( h* (i| j) k)"))
	assert.NoError(t, err)
	exp, err := p.Parse()
	assert.NoError(t, err)
	fmt.Printf("%v", exp)
}
