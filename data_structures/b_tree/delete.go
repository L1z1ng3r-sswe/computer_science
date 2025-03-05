package b_tree

import "fmt"

func (t *BTree) Delete(key int) {
	if len(t.Root.Keys) == 0 {
		return
	}

	t.Root.delete(key, t.T)

	if len(t.Root.Keys) == 0 && !t.Root.isLeaf() {
		t.Root = t.Root.Children[0]
	}
}

func (n *BTreeNode) delete(key int, t int) {
	i := 0
	for i < len(n.Keys) && n.Keys[i] < key {
		i++
	}

	if i < len(n.Keys) && n.Keys[i] == key { // Found the key
		if n.isLeaf() { // Have no children
			n.Keys = append(n.Keys[:i], n.Keys[i+1:]...)
		} else {
			if len(n.Children[i].Keys) >= t { // Replace with left child
				leftLargest := n.Children[i].largest()
				n.Keys[i] = leftLargest
				n.Children[i].delete(leftLargest, t)

			} else if len(n.Children[i+1].Keys) >= t { // Replace with right child
				rightSmallest := n.Children[i+1].smallest()
				n.Keys[i] = rightSmallest
				n.Children[i+1].delete(rightSmallest, t)

			} else { // Replace with right child
				n.merge(i)
				n.Children[i].delete(key, t)

			}
		}

		return
	}

	if n.isLeaf() {
		fmt.Println("Key is not found")
		return
	}

	n.Children[i].delete(key, t)
	if len(n.Children[i].Keys) < t {
		n.fillChild(i, t)
	}
}

func (n *BTreeNode) largest() int {
	curr := n

	for !curr.isLeaf() {
		curr = curr.Children[len(curr.Children)-1]
	}

	return curr.Keys[len(curr.Keys)-1]
}

func (n *BTreeNode) smallest() int {
	curr := n

	for !curr.isLeaf() {
		curr = curr.Children[0]
	}

	return curr.Keys[0]
}

func (n *BTreeNode) merge(i int) {
	left, right := n.Children[i], n.Children[i+1]

	left.Keys = append(left.Keys, n.Keys[i])     // Move parrent down
	left.Keys = append(left.Keys, right.Keys...) // Merge keys from right into left

	n.Keys = append(n.Keys[:i], n.Keys[i+1:]...)               // Remove first copy parrent
	n.Children = append(n.Children[:i+1], n.Children[i+2:]...) // Remove right children

	if !left.isLeaf() {
		left.Children = append(left.Children, right.Children...) // Merge right children into left children
	}
}

func (n *BTreeNode) fillChild(i int, t int) {
	if i > 0 && len(n.Children[i-1].Keys) > t { // Borrow from the left sibling
		n.borrowFromLeft(i)
	} else if i < len(n.Children)-1 && len(n.Children[i+1].Keys) > t { // Borrow from the right sibling
		n.borrowFromRight(i)
	} else { // Both left and right reached minimum -> merge
		if i > 0 {
			n.merge(i - 1)
		} else {
			n.merge(i)
		}
	}
}

func (n *BTreeNode) borrowFromLeft(i int) { // Works only on leaf nodes, use borrow from the children if u got children
	left, underFlow := n.Children[i-1], n.Children[i]

	// Borrow keys
	underFlow.Keys = underFlow.Keys[:len(underFlow.Keys)+1]
	copy(underFlow.Keys[1:], underFlow.Keys[0:])
	underFlow.Keys[0] = left.Keys[len(left.Keys)-1]
	left.Keys = left.Keys[:len(left.Keys)-1]
}

func (n *BTreeNode) borrowFromRight(i int) { // Works only on leaf nodes, use borrow from the children if u got children
	underFlow, right := n.Children[i], n.Children[i+1]

	// Borrow keys
	underFlow.Keys = append(underFlow.Keys, right.Keys[0])
	copy(right.Keys[0:], right.Keys[1:])
	right.Keys = right.Keys[:len(right.Keys)-1]
}
