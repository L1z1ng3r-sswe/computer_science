package b_tree

func (t *BTree) Insert(key int) {
	if len(t.Root.Nodes) == 2*t.T {
		newRoot := newNode(t.T)
		newRoot.Children = append(newRoot.Children, t.Root)

		newRoot.splitChild(0, t.T)
		t.Root = newRoot
	}

	t.Root.insert(key, t.T)
}

func (n *BTreeNode) insert(key int, t int) {
	i := len(n.Nodes) - 1

	if n.isLeaf() {
		n.Nodes = n.Nodes[:len(n.Nodes)+1] // Expand

		for i >= 0 && n.Nodes[i] > key {
			n.Nodes[i+1] = n.Nodes[i]
			i--
		}

		n.Nodes[i+1] = key
	} else {
		for i >= 0 && n.Nodes[i] > key {
			i--
		}
		i++

		if len(n.Children[i].Nodes) == 2*t { // Child is full
			n.splitChild(i, t)

			if key > n.Nodes[i] {
				i++
			}
		}

		n.Children[i].insert(key, t)
	}
}
func (n *BTreeNode) splitChild(i int, t int) { // i - child in children that is full
	fullNode := n.Children[i]

	newNode := newNode(t)
	newNode.Nodes = append(newNode.Nodes, fullNode.Nodes[t+1:]...)
	fullNode.Nodes = fullNode.Nodes[:t]

	if !fullNode.isLeaf() {
		newNode.Children = append(newNode.Children, fullNode.Children[t+1:]...)
		fullNode.Children = fullNode.Children[:t+1]
	}

	n.Nodes = n.Nodes[:len(n.Nodes)+1] // Expand slice
	copy(n.Nodes[i+1:], n.Nodes[i:])
	n.Nodes[i] = fullNode.Nodes[t]

	n.Children = n.Children[:len(n.Children)+1]
	copy(n.Children[i+2:], n.Children[i+1:])
	n.Children[i+1] = newNode
}
