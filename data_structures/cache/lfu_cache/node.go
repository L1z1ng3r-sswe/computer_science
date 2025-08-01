package lfu_cache

type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
	Freq  int
}

// default freq is 1
func NewNode(key int, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
		Freq:  1,
	}
}

func (node *Node) remove() {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}
