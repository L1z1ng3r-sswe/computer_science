package b_tree

func (t *BTree) Search(key int) bool {
	return t.Root.search(key)
}

func (n *BTreeNode) search(key int) bool {
	i := 0

	for i < len(n.Nodes) && n.Nodes[i] < key {
		i++
	}

	if i < len(n.Nodes) && n.Nodes[i] == key {
		return true
	}

	if n.isLeaf() {
		return false
	}

	return n.Children[i].search(key)
}
