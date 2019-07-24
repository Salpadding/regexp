package fsa

import "errors"

type node struct {
	token      *token
	children   []*node
	leftChild  *node
	rightChild *node
}

func parse(tokens []*token) *node {
	var pc int
	n := parseMulti(tokens, &pc, nil)
	addModifier(n)
	return refactor(n)
}

// parse tokens as multi-fork tree
func parseMulti(tokens []*token, pc *int, left *node) *node {
	if left == nil {
		left = &node{
			token:    concat,
			children: []*node{},
		}
	}

	for *pc < len(tokens) {
		tk := tokens[*pc]
		*pc ++
		switch tk.code {
		case tokenLeftParentheses:
			left.children = append(left.children, parseMulti(tokens, pc, nil))
		case tokenRightParentheses:
			return left
		case tokenOr:
			tmp := &node{
				token:    tk,
				children: []*node{left},
			}
			tmp.children = append(tmp.children, parseMulti(tokens, pc, nil))
			return tmp
		default:
			left.children = append(left.children, &node{
				token: tk,
			})
		}
	}
	return left
}

// add precedence *, ?, +
func addModifier(tree *node) {
	if tree.children == nil {
		return
	}
	var tmp []*node
	for i := 0; i < len(tree.children); i++ {
		if i+1 == len(tree.children) {
			tmp = append(tmp, tree.children[i])
			break
		}
		switch tree.children[i+1].token.code {
		case tokenClosure, tokenOneOrMore, tokenNoneOrOne:
			tmp = append(tmp, &node{
				token:     tree.children[i+1].token,
				leftChild: tree.children[i],
			})
			i++
		default:
			tmp = append(tmp, tree.children[i])
		}

	}
	tree.children = tmp
	for _, n := range tree.children {
		addModifier(n)
	}
}

// convert multi-fork tree to binary tree
func refactor(tree *node) *node {
	switch tree.token.code {
	case tokenOr:
		return &node{
			token:      tree.token,
			leftChild:  refactor(tree.children[0]),
			rightChild: refactor(tree.children[1]),
		}
	case tokenConcat:
		if tree.children == nil || len(tree.children) == 0 {
			return tree
		}
		if len(tree.children) == 1 {
			return refactor(tree.children[0])
		}
		return &node{
			token:     concat,
			leftChild: refactor(tree.children[0]),
			rightChild: refactor(&node{
				token:    concat,
				children: tree.children[1:],
			}),
		}
	case tokenClosure, tokenOneOrMore, tokenNoneOrOne:
		tree.leftChild = refactor(tree.leftChild)
		return tree
	default:
		return tree
	}
}

func validateTree(node *node) (bool, error) {
	err := errors.New("invalid input regexp")
	if node == nil || node.children != nil {
		return false, err
	}
	switch node.token.code {
	case tokenConcat, tokenOr:
		if node.leftChild == nil || node.rightChild == nil {
			return false, err
		}
		ok, _ := validateTree(node.leftChild)
		if !ok {
			return ok, err
		}
		ok, _ = validateTree(node.rightChild)
		if !ok {
			return ok, err
		}
	case tokenClosure, tokenOneOrMore, tokenNoneOrOne:
		if node.leftChild == nil || node.rightChild != nil {
			return false, err
		}
		ok, _ := validateTree(node.leftChild)
		if !ok {
			return ok, err
		}
	}
	return true, nil
}
