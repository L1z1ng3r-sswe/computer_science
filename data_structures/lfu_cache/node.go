package lfu_cache

type Node struct {
	Value      int
	Freq       int
	Next, Prev *Node
}

// default value of freq is 1, next and prev are nil
func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Freq:  1,
		Next:  nil,
		Prev:  nil,
	}
}

func (n *Node) removeNode() {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}
