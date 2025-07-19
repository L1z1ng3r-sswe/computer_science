package lfu_cache

type LRUCache struct {
	Head, Tail *Node
}

func NewLRUCache() *LRUCache {
	head, tail := NewNode(0), NewNode(0)
	tail.Prev = head
	head.Next = tail

	return &LRUCache{
		Head: head,
		Tail: tail,
	}
}

func (c *LRUCache) isEmpty() bool {
	if c.Head.Next == c.Tail {
		return true
	}
	return false
}

func (c *LRUCache) insert(node *Node) {
	node.Next = c.Head.Next
	node.Prev = c.Head
	c.Head.Next.Prev = node
	c.Head.Next = node
}
