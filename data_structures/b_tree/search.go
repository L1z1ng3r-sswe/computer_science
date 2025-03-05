package b_tree

import "fmt"

func (t *BTree) Search(key int) bool {
	return t.Root.search(key)
}

func (n *BTreeNode) search(key int) bool {
	i := 0
	for i < len(n.Keys) && key > n.Keys[i] { // Linear search, but i was considering a check - if T > 100 than use the binary search but decided to omit
		i++
	}

	if i < len(n.Keys) && key == n.Keys[i] {
		return true
	}

	if n.isLeaf() {
		fmt.Println("Key is not found")
		return false
	}

	return n.Children[i].search(key)
}
