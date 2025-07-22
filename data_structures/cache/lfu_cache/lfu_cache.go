package lfu_cache

type LFUCache struct {
	Capacity int
	Elements map[int]*Node
	FreqMap  map[int]*LRUCache
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Capacity: capacity,
		Elements: make(map[int]*Node, capacity),
		FreqMap:  map[int]*LRUCache{1: NewLRUCache()},
	}
}

func (c *LFUCache) Get(key int) int {
	if node, ok := c.Elements[key]; ok {
		c.updateFreq(node, node.Freq, node.Freq+1)
		return node.Value
	}

	return -1
}

func (c *LFUCache) Put(key int, value int) {
	if node, ok := c.Elements[key]; ok { // update
		node.Value = value
		c.updateFreq(node, node.Freq, node.Freq+1)
	} else { // create
		node = NewNode(value)
		lru := c.FreqMap[1]
		lru.insert(node)
	}
}

// Moves node to new freqMap level
// after removing node we need to verify
// that the old lru is not empty.
// if it is so - delete it from FreqMap
func (c *LFUCache) updateFreq(node *Node, oldFreq, newFreq int) {
	node.Freq = newFreq
	node.removeNode()

	if c.FreqMap[oldFreq].isEmpty() {
		delete(c.FreqMap, oldFreq)
	}

	lru := c.getOrCreateFreqLRU(newFreq)
	lru.insert(node)
}

// Inserts new lru in FreqMap if not created yet, or return if created
func (c *LFUCache) getOrCreateFreqLRU(freq int) *LRUCache {
	lru := c.FreqMap[freq]
	if lru == nil {
		return c.newFreqLRU(freq)
	}

	return lru
}

// Inserts new freq and returns lru of this freq
func (c *LFUCache) newFreqLRU(freq int) *LRUCache {
	lru := NewLRUCache()
	c.FreqMap[freq] = lru
	return lru
}
