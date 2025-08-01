package lfu_cache

type LRUCache struct {
	Head, Tail *Node
}

func NewLRUCache() *LRUCache {
	head, tail := NewNode(0, 0), NewNode(0, 0)
	head.Next = tail
	tail.Prev = head

	return &LRUCache{
		Head: head,
		Tail: tail,
	}
}

// check if lru is empty
func (c *LRUCache) isEmpty() bool {
	return c.Head.Next == c.Tail
}

func (c *LRUCache) addToFront(node *Node) {
	node.Next = c.Head.Next
	node.Prev = c.Head
	c.Head.Next.Prev = node
	c.Head.Next = node
}

// works only when lru not empty, check is empty before
func (c *LRUCache) popLRU() *Node {
	lruNode := c.Tail.Prev

	lruNode.remove()

	return lruNode
}
