package binary_search_tree

import "fmt"

type BST struct {
	Val       int
	LeftNode  *BST
	RightNode *BST
}

func (b *BST) Add(val int) {
	if b == nil {
		return
	}

	if b.Val > val {
		if b.LeftNode == nil {
			b.LeftNode = &BST{Val: val}
		} else {
			b.LeftNode.Add(val)
		}
	} else {
		if b.RightNode == nil {
			b.RightNode = &BST{Val: val}
		} else {
			b.RightNode.Add(val)
		}
	}
}

func (b *BST) Delete(val int) *BST {
	if b == nil {
		return nil
	}

	if b.Val > val {
		b.LeftNode = b.LeftNode.Delete(val)
	} else if b.Val < val {
		b.RightNode = b.RightNode.Delete(val)
	} else {
		if b.LeftNode == nil {
			return b.RightNode
		} else if b.RightNode == nil {
			return b.LeftNode
		}

		minRight := MinNode(b.RightNode)
		b.Val = minRight.Val
		b.RightNode = b.RightNode.Delete(minRight.Val)
	}

	return b
}

func (b *BST) Search(val int) *BST {
	if b == nil {
		return nil
	}

	if b.Val > val {
		return b.LeftNode.Search(val)
	} else if b.Val < val {
		return b.RightNode.Search(val)
	}
	return b
}

func (b *BST) PrettyVisualise(lvl int) string {
	if b == nil {
		return ""
	}

	left := b.LeftNode.PrettyVisualise(lvl + 1)
	head := fmt.Sprintf("%s%d\n", Indent(lvl), b.Val)
	right := b.RightNode.PrettyVisualise(lvl + 1)

	return right + head + left
}

func Indent(lvl int) string {
	indent := ""

	for i := 0; i < lvl; i++ {
		indent += " "
	}

	return indent
}

func MinNode(node *BST) *BST {
	if node == nil {
		return node
	}

	curr := node

	for curr.LeftNode != nil {
		curr = curr.LeftNode
	}

	return curr
}
