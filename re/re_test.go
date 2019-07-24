package re

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchEmail(t *testing.T) {
	re, err := Compile(`(\d|\w)+@(\d|\w)+\.(\d|\w)+`)
	assert.NoError(t, err)
	re.InputString("m6567fc@outlook.com")
	assert.True(t, re.IsAccept())
	re.Reset()
	re.InputString("abbbbb@yyy")
	assert.False(t, re.IsAccept())
}

func TestMatchEmail2(t *testing.T) {
	re, err := Compile(`[a-z0-9A-Z]+@[a-z0-9A-Z]+\.[a-z0-9A-Z]+`)
	assert.NoError(t, err)
	re.InputString("m6567fc@outlook.com")
	assert.True(t, re.IsAccept())
	re.Reset()
	re.InputString("abbbbb@yyy")
	assert.False(t, re.IsAccept())
}

func TestMatchFloat(t *testing.T) {
	re, err := Compile(`-?[0-9]+(\.[0-9]+)?`)
	assert.NoError(t, err)
	assert.True(t, re.Match("0.09"))
	assert.False(t, re.Match("NaN"))
}

func TestMatchHex(t *testing.T) {
	re, err := Compile(`0x[a-f0-9]+`)
	assert.NoError(t, err)
	assert.True(t, re.Match("0xff"))
	assert.False(t, re.Match("0xg"))
}

func TestMatchKeyword(t *testing.T){
	re, err := Compile(`break|default|func|interface|select|case|defer|go|map|struct|chan|else|goto|package|switch|const|fallthrough|if|range|type|continue|for|import|return|var`)
	assert.NoError(t, err)
	assert.True(t, re.Match("break"))
	assert.True(t, re.Match("default"))
	assert.True(t, re.Match("func"))
}