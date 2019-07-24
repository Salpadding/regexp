package fsa

import "errors"

type node struct {
	token      *token
	children   []*node
	leftChild  *node
	rightChild *node
}

func parse(tokens []*token) (*node, error) {
	var pc int
	n := parseMulti(tokens, &pc, nil)
	if err := addModifier(n); err != nil {
		return nil, err
	}
	return refactor(n), nil
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
func addModifier(tree *node) error {
	if tree.children == nil {
		return nil
	}
	var (
		tmp []*node
		err error
	)
	for _, n := range tree.children {
		switch n.token.code {
		// *, +, ?
		case tokenClosure, tokenOneOrMore, tokenNoneOrOne:
			if tmp == nil || len(tmp) < 1 {
				return errors.New("unexpected modifier")
			}
			l := len(tmp)
			tmp[l-1] = &node{
				token:     n.token,
				leftChild: tmp[l-1],
			}
		default:
			err = addModifier(n)
			if err != nil {
				return err
			}
			tmp = append(tmp, n)
		}

	}
	tree.children = tmp
	return nil
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
