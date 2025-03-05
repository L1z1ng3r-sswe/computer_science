type LRUCache struct {
	cache      map[int]*Node
	head, rear *Node
	capacity   int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache:    make(map[int]*Node, capacity),
		capacity: capacity,
		head:     &Node{},
		rear:     &Node{},
	}

	lru.head.next = lru.rear
	lru.rear.prev = lru.head.next

	return lru
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		node.cut()
		c.insert(node)

		return node.val
	}

	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.cache[key]; ok { // Just replace
		node.cut()
		c.insert(node)

		node.val = value
	} else {
		if len(c.cache) == c.capacity { // Cache is full - shrink
			delete(c.cache, c.rear.prev.key)

			c.rear.prev.prev.next = c.rear
			c.rear.prev = c.rear.prev.prev
		}

		node := &Node{
			key: key,
			val: value,
		}

		c.insert(node)

		c.cache[key] = node
	}
}

func (c *LRUCache) insert(node *Node) {
	node.next = c.head.next
	node.prev = c.head
	c.head.next.prev = node
	c.head.next = node
}

type Node struct {
	key, val   int
	prev, next *Node
}

func (n *Node) cut() {
	n.next.prev = n.prev
	n.prev.next = n.next
}