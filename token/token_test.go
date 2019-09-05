package token

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

type testEnum byte

const (
	_ testEnum = iota
	testEnum1
	testEnum2
)

type testType interface {
	testType()
}

type testType2 interface {
	testEnum() testEnum
}

type tp1 string

func (tp1) testType() {}

func (tp1) testEnum() testEnum {
	return testEnum1
}

type tp2 string

func (tp2 tp2) testEnum() testEnum {
	return testEnum2
}

func (tp2) testType() {}

func istp1(t testType) bool {
	_, ok := t.(tp1)
	return ok
}

func istp2(t testType2) bool {
	return t.testEnum() == testEnum2
}

// 0.25 ns/op
func Benchmark1(b *testing.B) {
	for i := 0; i < math.MaxInt32; i++{
		istp1(tp1("====="))
	}
}

// 3589436900 ns/op
func Benchmark2(b *testing.B) {
	for i:= 0; i < math.MaxInt32; i++{
		istp2(tp2("====="))
	}
}

func Test(t *testing.T){
	assert.True(t, istp1(tp1("")))
	assert.True(t, istp1(tp1("2")))
	assert.False(t, istp1(tp2("")))
}
