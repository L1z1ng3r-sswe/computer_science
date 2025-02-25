package b_tree

import "fmt"

func (t *BTree) Delete(key int) {
	t.Root.delete(key, t.T)
}

func (n *BTreeNode) delete(key int, t int) {
	i := 0

	for i < len(n.Nodes) && key > n.Nodes[i] {
		i++
	}

	if i < len(n.Nodes) && n.Nodes[i] == key { // Found the key
		if len(n.Children) == 0 { // Leaf node
			n.Nodes = append(n.Nodes[t:], n.Nodes[t+1:]...)
		} else {
			if len(n.Children[i].Nodes) > t { // Left got enough
				leftGreatest := n.Children[i].largest()
				n.Nodes[i] = leftGreatest
				n.Children[i].delete(leftGreatest, t)
			} else if len(n.Children[i+1].Nodes) > t { // Right got enough
				rightSmallest := n.Children[i+1].smallest()
				n.Nodes[i] = rightSmallest
				n.Children[i+1].delete(rightSmallest, t)
			} else { // both have not enough -> merge them
				n.mergeNodes(i)
				n.Children[i].delete(i, t)
			}
		}
		return
	}

	if len(n.Children) == 0 { // Leaf node
		fmt.Println("Key is not found")
		return
	}

	n.Children[i].delete(key, t)

	if len(n.Children[i].Nodes) < t {
		n.fillNode(i, t)
	}
}

func (n *BTreeNode) largest() int {
	curr := n

	for !curr.isLeaf() {
		curr = curr.Children[len(curr.Children)-1]
	}

	return curr.Nodes[len(curr.Nodes)-1]
}

func (n *BTreeNode) smallest() int {
	curr := n

	for !curr.isLeaf() {
		curr = curr.Children[0]
	}

	return curr.Nodes[0]
}

func (n *BTreeNode) mergeNodes(i int) {
	leftNode, rightNode := n.Children[i], n.Children[i+1]

	leftNode.Nodes = append(leftNode.Nodes, n.Nodes[i])         // Merge keys
	leftNode.Nodes = append(leftNode.Nodes, rightNode.Nodes...) // Merge keys

	n.Nodes = append(n.Nodes[:i], n.Nodes[i+1:]...)

	if len(leftNode.Children) > 0 { // Non leaf node
		leftNode.Children = append(leftNode.Children, rightNode.Children...) // Merge children
		n.Children = append(n.Children[:i+1], n.Children[i+2:]...)           // Delete left node
	}
}

func (n *BTreeNode) fillNode(i int, t int) {
	if i > 0 && len(n.Children[i-1].Nodes) > t { // left sibling got enough, borrow
		n.borrowFromLeft(i)
	} else if i < len(n.Children)-1 && len(n.Children[i+1].Nodes) > t { // right sibling got enough, borrow
		n.borrowFromRight(i)
	} else { // both sibling are under flow -> merge
		if i > 0 {
			n.mergeNodes(i - 1)
		} else {
			n.mergeNodes(i)
		}
	}
}

func (n *BTreeNode) borrowFromLeft(i int) {
	left, underFlow := n.Children[i-1], n.Children[i]

	underFlow.Nodes = underFlow.Nodes[:len(underFlow.Nodes)+1]
	copy(underFlow.Nodes[1:], underFlow.Nodes[0:])
	underFlow.Nodes[0] = n.Nodes[i-1]

	n.Nodes[i-1] = left.Nodes[len(left.Nodes)-1]
	left.Nodes = left.Nodes[:len(left.Nodes)-1]

	if len(left.Children) > 0 {
		underFlow.Children = underFlow.Children[:len(underFlow.Children)+1]
		copy(underFlow.Children[1:], underFlow.Children[0:])
		underFlow.Children[0] = left.Children[len(left.Children)-1]

		left.Children = left.Children[:len(left.Children)-1]
	}
}

func (n *BTreeNode) borrowFromRight(i int) {
	underFlow, right := n.Children[i], n.Children[i+1]

	underFlow.Nodes = append(underFlow.Nodes, n.Nodes[i])
	n.Nodes[i] = right.Nodes[0]
	copy(right.Nodes[0:], right.Nodes[1:])
	right.Nodes = right.Nodes[:len(right.Nodes)-1]

	if len(right.Children) > 0 {
		underFlow.Children = append(underFlow.Children, right.Children[0])
		copy(right.Children[0:], right.Children[1:])
		right.Children = right.Children[:len(right.Children)-1]
	}
}
