package re

import (
	"github.com/stretchr/testify/assert"
	regexp2 "regexp"
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

func TestMatchKeyword(t *testing.T) {
	re, err := Compile(`break|default|func|interface|select|case|defer|go|map|struct|chan|else|goto|package|switch|const|fallthrough|if|range|type|continue|for|import|return|var`)
	assert.NoError(t, err)
	assert.True(t, re.Match("break"))
	assert.True(t, re.Match("default"))
	assert.True(t, re.Match("func"))
}

func TestWisdomURL(t *testing.T) {
	re, err := Compile(`wisdom://([0-9a-f]+@)?((\d+\.\d+\.\d+\.\d+)|[0-9a-zA-Z]+)(:[0-9]+)?`)
	assert.NoError(t, err)
	assert.True(t, re.Match("wisdom://76a3f5787062ffd12425b27e14f29348a7407b42ebfdff8e14543e6356e10530@192.168.0.104:9005"))
	assert.True(t, re.Match("wisdom://192.168.0.104:9005"))
	assert.True(t, re.Match("wisdom://localhost"))
	assert.False(t, re.Match("wisdom://76a3f5787062ffd12425b27e14f29348a7407b42ebfdff8e14543e6356e10530@192.168.0:9005"))
}

func Benchmark1(b *testing.B) {
	re, err := Compile(`wisdom://([0-9a-f]+@)?((\d+\.\d+\.\d+\.\d+)|[0-9a-zA-Z]+)(:[0-9]+)?`)
	if err != nil{
		b.Fail()
	}
	b.StartTimer()
	for i := 0; i < 100000; i++{
		re.Match("wisdom://76a3f5787062ffd12425b27e14f29348a7407b42ebfdff8e14543e6356e10530@192.168.0.104:9005")
	}
	b.StopTimer()
}

func Benchmark2(b *testing.B) {
	re, err := regexp2.Compile(`wisdom://([0-9a-f]+@)?((\d+\.\d+\.\d+\.\d+)|[0-9a-zA-Z]+)(:[0-9]+)?`)
	if err != nil{
		b.Fail()
	}
	b.StartTimer()
	for i := 0; i < 100000; i++{
		re.MatchString("wisdom://76a3f5787062ffd12425b27e14f29348a7407b42ebfdff8e14543e6356e10530@192.168.0.104:9005")
	}
	b.StopTimer()
}