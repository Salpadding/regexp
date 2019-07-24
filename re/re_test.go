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
}
