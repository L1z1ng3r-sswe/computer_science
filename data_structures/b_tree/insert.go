package b_tree

func (t *BTree) Insert(key int) {
	if len(t.Root.Keys) == 2*t.T+1 {
		newRoot := newBTreeNode(t.T)
		newRoot.Children = append(newRoot.Children, t.Root)
		newRoot.split(0, t.T)

		t.Root = newRoot
	}

	t.Root.insert(key, t.T-1)
}

func (n *BTreeNode) insert(key int, t int) {
	i := len(n.Keys) - 1

	if n.isLeaf() {
		n.Keys = n.Keys[:len(n.Keys)+1]

		for i >= 0 && key < n.Keys[i] {
			n.Keys[i+1] = n.Keys[i]
			i--
		}

		n.Keys[i+1] = key
	} else {
		for i >= 0 && key < n.Keys[i] {
			i--
		}

		i++

		if len(n.Children[i].Keys) == 2*t+1 {
			n.split(i, t)
			if key > n.Keys[i] {
				i++
			}
		}

		n.Children[i].insert(key, t)
	}
}

func (n *BTreeNode) split(i int, t int) {
	fullChild := n.Children[i]

	n.Keys = n.Keys[:len(n.Keys)+1]
	if i+1 < len(n.Keys) { // not a last element
		copy(n.Keys[i+1:], n.Keys[i:])
	}
	n.Keys[i] = fullChild.Keys[t]

	newNode := newBTreeNode(t)
	newNode.Keys = fullChild.Keys[t+1:]
	fullChild.Keys = fullChild.Keys[:t]

	if !fullChild.isLeaf() {
		newNode.Children = fullChild.Children[t:]
		fullChild.Children = fullChild.Children[:t]
	}

	n.Children = n.Children[:len(n.Children)+1]
	if i+2 < len(n.Children) { // not a last element
		copy(n.Children[i+2:], n.Children[i+1:])
	}
	n.Children[i+1] = newNode
}
