package b_tree

import "fmt"

type BTree struct {
	T    int // minimum number of keys per each Node
	Root *BTreeNode
}

func New(t int) *BTree {
	if t < 1 {
		fmt.Println("t must be more than 0")
		return nil
	}

	return &BTree{
		T:    t,
		Root: newBTreeNode(t),
	}
}

type BTreeNode struct {
	Keys     []int        // 2*T + 1 - max number of Keys per Node
	Children []*BTreeNode // 2*T + 2 - max number of children per Node
}

func newBTreeNode(t int) *BTreeNode {
	return &BTreeNode{
		Keys:     make([]int, 0, t*2+1),
		Children: make([]*BTreeNode, 0, t*2+2),
	}
}

func (n *BTreeNode) isLeaf() bool {
	return len(n.Children) == 0
}
