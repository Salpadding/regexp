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