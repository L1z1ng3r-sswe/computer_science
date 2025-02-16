package b_tree

import "fmt"

type BTree struct { // BTree serves as a "dummy" node in case of full root node
	Root *BTreeNode
	T    int // t*2 - max number of nodes, t*2 + 1 - max number of children
}

type BTreeNode struct {
	Nodes    []int
	Children []*BTreeNode
}

func NewBTree(t int) *BTree { // t must be more than 1
	if t < 2 {
		fmt.Println("t is less than 2")
		return nil
	}

	return &BTree{
		Root: newNode(t),
		T:    t,
	}
}

func (n *BTreeNode) isLeaf() bool {
	return len(n.Children) == 0
}

func newNode(t int) *BTreeNode {
	return &BTreeNode{
		Nodes:    make([]int, 0, t*2),
		Children: make([]*BTreeNode, 0, t*2+1),
	}
}
