package json

import "github.com/Salpadding/regexp/re"

// json tokenizer

var (
	stringMatcher   = re.Compile(`".*"`)
	numberMatcher   = re.Compile(`(-?)\d*(\.\d*)?`)
	boolMatcher     = re.Compile("(true)|(false)")
	reversedMatcher = re.Compile("{|}|,|[|]")
)
