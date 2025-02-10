package b_tree

import "fmt"

type BTree struct {
	T    int
	Root *BTreeNode
}

type BTreeNode struct {
	Nodes    []int
	Children []*BTreeNode
	Leaf     bool
}

func New(T int) (*BTree, error) {
	if T < 1 {
		return nil, fmt.Errorf("T is less than 1")
	}

	return &BTree{
		T:    T,
		Root: &BTreeNode{Leaf: true},
	}, nil
}

func (t *BTree) Search(key int) int {
	return t.Root.search(key)
}

func (n *BTreeNode) search(key int) int {
	var i int
	for i < len(n.Nodes) && key > n.Nodes[i] {
		i++
	}

	if i < len(n.Nodes) && key == n.Nodes[i] {
		return key
	}

	if n.Leaf {
		return -1
	}

	return n.Children[i].search(key)
}

func (t *BTree) Insert(key int) {
	root := t.Root
	if len(root.Nodes) == 2*t.T {
		newRoot := &BTreeNode{
			Children: []*BTreeNode{root},
		}
		newRoot.splitChild(0, t.T)
		t.Root = newRoot
	}
	t.Root.insert(key, t.T)
}

func (n *BTreeNode) insert(key int, t int) {
	i := len(n.Nodes) - 1

	if n.Leaf {
		n.Nodes = append(n.Nodes, 0)
		for i >= 0 && key < n.Nodes[i] {
			n.Nodes[i+1] = n.Nodes[i]
			i--
		}
		n.Nodes[i+1] = key
		return
	}

	for i >= 0 && key < n.Nodes[i] {
		i--
	}

	i++

	if len(n.Children[i].Nodes) == 2*t {
		n.splitChild(i, t)
		if key > n.Nodes[i] {
			i++
		}
	}

	n.Children[i].insert(key, t)
}

func (n *BTreeNode) splitChild(i int, t int) {
	child := n.Children[i]
	newChild := &BTreeNode{
		Nodes: child.Nodes[t+1:],
		Leaf:  child.Leaf,
	}

	if !child.Leaf {
		newChild.Children = child.Children[t+1:]
		child.Children = child.Children[:t+1]
	}

	n.Nodes = append(n.Nodes[:i+1], append([]int{child.Nodes[t]}, n.Nodes[i:]...)...)
	n.Children = append(n.Children[:i+1], append([]*BTreeNode{newChild}, n.Children[i+1:]...)...)

	child.Nodes = child.Nodes[:t]
}
