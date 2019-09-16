package parser

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/Salpadding/regexp/token"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	p, err := New(bytes.NewBufferString(`(\d|\w)+@(\d|\w)+\.(\d|\w)+)`))
	assert.NoError(t, err)
	var tks []token.Token
	for {
		_, ok := p.current.(token.EOF)
		if ok {
			break
		}
		tks = append(tks, p.current)
		_, err := p.nextToken()
		assert.NoError(t, err)
	}
	fmt.Printf("%v", tks)
}

func Test(t *testing.T) {
	p, err := New(bytes.NewBufferString(`(\d|\w)+@(\d|\w)+\.(\d|\w)+)`))
	assert.NoError(t, err)
	exp, err := p.Parse()
	assert.NoError(t, err)
	fmt.Printf("%v", exp)
}

func Test2(t *testing.T) {
	p, err := New(bytes.NewBufferString(`[a-z0-9A-Z]+@[a-z0-9A-Z]+\.[a-z0-9A-Z]+`))
	assert.NoError(t, err)
	exp, err := p.Parse()
	assert.NoError(t, err)
	fmt.Printf("%v", exp)
}
