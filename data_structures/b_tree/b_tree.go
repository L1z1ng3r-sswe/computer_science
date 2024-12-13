package b_tree

import "errors"

// at most 2T of keys and 2T+1 of children
func New(T int) (*BTree, error) {
	if T < 1 {
		return nil, errors.New("T must be at least 1")
	}

	return &BTree{
		T:    T,
		Root: &BTreeNode{},
	}, nil
}

type BTree struct {
	Root *BTreeNode
	T    int // minimum number of nodes
}

// Insert adds a key to the B-tree.
func (t *BTree) Insert(key int) {
	if len(t.Root.Nodes) >= t.T*2 {
		// root is full, split it

		newRoot := &BTreeNode{
			Nodes:    make([]int, 0, t.T*2),
			Children: append(make([]*BTreeNode, 0, t.T*2+1), t.Root),
		}

		newRoot.splitChild(0, t.T)
		t.Root = newRoot
	}

	t.Root.insert(key, t.T)
}

func (t *BTree) Search(key int) (*BTreeNode, bool) {
	return t.Root.search(key)
}

type BTreeNode struct {
	Nodes    []int        // Keys in the node (max: T*2)
	Children []*BTreeNode // Child pointers (max: T*2 + 1)
}

func (n *BTreeNode) search(key int) (*BTreeNode, bool) {
	currNode := n

	for {
		i := 0

		for i < len(currNode.Nodes) && key > currNode.Nodes[i] {
			i++
		}

		if i < len(currNode.Nodes) && key == currNode.Nodes[i] {
			return currNode, true
		}

		if len(currNode.Children) == 0 {
			return nil, false
		}

		currNode = currNode.Children[i]
	}
}

func (n *BTreeNode) insert(key int, T int) {
	currNode := n

	for {
		i := len(currNode.Nodes) - 1

		if len(currNode.Children) == 0 {
			// it is a leaf, insert the key in the current node
			currNode.Nodes = append(currNode.Nodes, 0) // Increase length by 1
			// currNode.Nodes = currNode.Nodes[:len(currNode.Nodes)+1] this entry works the same way, but it is unsafe

			for i >= 0 && key < currNode.Nodes[i] {
				currNode.Nodes[i+1] = currNode.Nodes[i]
				i--
			}
			currNode.Nodes[i+1] = key
			return
		}

		for i >= 0 && key < currNode.Nodes[i] {
			i--
		}
		i++

		if len(currNode.Children[i].Nodes) >= 2*T {
			// child is full, split it

			currNode.splitChild(i, T)
			if key > currNode.Nodes[i] {
				i++
			}
		}
		currNode = currNode.Children[i]
	}
}

func (n *BTreeNode) splitChild(i int, T int) {
	fullChild := n.Children[i]

	newNode := &BTreeNode{
		Nodes:    append(make([]int, 0, 2*T), fullChild.Nodes[T+1:]...),
		Children: append(make([]*BTreeNode, 0, 2*T+1), fullChild.Children[T+1:]...),
	}

	fullChild.Nodes = fullChild.Nodes[:T]
	fullChild.Children = fullChild.Children[:T+1]

	n.Nodes = append(n.Nodes[:i], append([]int{fullChild.Nodes[T]}, n.Nodes[i:]...)...)
	n.Children = append(n.Children[:i+1], append([]*BTreeNode{newNode}, n.Children[i+1:]...)...)
}
