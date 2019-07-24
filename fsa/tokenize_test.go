package fsa

//func TestTokenize(t *testing.T) {
//	res := tokenize(bytes.NewBufferString("( a | b ) * cd"))
//	assert.NotEmpty(t, res)
//	assert.Equal(t, 10, len(res))
//	assert.Equal(t, cache['('], res[0])
//	assert.Equal(t, tokenChar, res[1].code)
//	assert.Equal(t, tokenOr, res[2].code)
//	assert.Equal(t, tokenChar, res[3].code)
//	assert.Equal(t, cache[')'], res[4])
//	assert.Equal(t, cache['*'], res[5])
//	assert.Equal(t, concat, res[6])
//	assert.Equal(t, tokenChar, res[7].code)
//	assert.Equal(t, concat, res[8])
//	assert.Equal(t, tokenChar, res[9].code)
//}
//
//func TestGetRune(t *testing.T){
//	assert.True(t, '中' < 0xffff)
//	assert.True(t, '国' > 0xff)
//	fmt.Println(string(rune(65)))
//}