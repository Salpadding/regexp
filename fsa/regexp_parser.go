package fsa

func New(regexp string) (*NFA, error) {
	tokens, err := tokenize(regexp)
	if err != nil {
		return nil, err
	}
	node := parse(tokens)
	ok, err := validateTree(node)
	if !ok {
		return nil, err
	}
	return eval(node), nil
}

func eval(tree *node) *NFA {
	switch tree.token.code {
	case tokenConcat:
		return eval(tree.leftChild).concat(eval(tree.rightChild))
	case tokenOr:
		return eval(tree.leftChild).or(eval(tree.rightChild))
	case tokenClosure:
		return eval(tree.leftChild).kleen()
	case tokenNoneOrOne:
		return eval(tree.leftChild).noneOrOne()
	case tokenOneOrMore:
		return eval(tree.leftChild).oneOrMore()
	case tokenChar:
		return NewChar(tree.token.value)
	case tokenDigital:
		return newDigital()
	case tokenLetters:
		return newLetters()
	case tokenWildcard:
		return newWildCard()
	case tokenNonDigital:
		return newNonDigital()
	case tokenNonLetter:
		return newNonLetter()
	}
	return nil
}
