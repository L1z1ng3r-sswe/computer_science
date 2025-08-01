package lfu_cache

type LFUCache struct {
	Capacity int
	MinFreq  int
	Nodes    map[int]*Node
	FreqMap  map[int]*LRUCache // key is frequency
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Capacity: capacity,
		MinFreq:  1,
		Nodes:    make(map[int]*Node),
		FreqMap:  make(map[int]*LRUCache),
	}
}

func (c *LFUCache) Get(key int) int {
	if node, ok := c.Nodes[key]; ok {
		c.increaseFreq(node)
		return node.Value
	}

	return -1
}

func (c *LFUCache) Put(key int, value int) {
	if node, ok := c.Nodes[key]; ok { // update
		node.Value = value
		c.increaseFreq(node)
	} else { // push
		if len(c.Nodes) == c.Capacity { // reach limit
			victimLRU := c.FreqMap[c.MinFreq]
			victimNode := victimLRU.popLRU()
			delete(c.Nodes, victimNode.Key)

			if victimLRU.isEmpty() {
				delete(c.FreqMap, c.MinFreq)

				if victimLRU.isEmpty() {
					delete(c.FreqMap, c.MinFreq)
				}
			}
		}

		c.MinFreq = 1
		newNode := NewNode(key, value)

		c.Nodes[key] = newNode
		c.ensureLRU(c.MinFreq).addToFront(newNode)
	}
}

func (c *LFUCache) increaseFreq(node *Node) {
	oldFreq, newFreq := node.Freq, node.Freq+1
	node.Freq = newFreq

	// remove from old lru
	node.remove()
	if c.FreqMap[oldFreq].isEmpty() {
		delete(c.FreqMap, oldFreq)
		if oldFreq == c.MinFreq {
			c.MinFreq = newFreq
		}
	}

	// insert to new lru
	c.ensureLRU(newFreq).addToFront(node)
}

func (c *LFUCache) ensureLRU(freq int) *LRUCache {
	cache := c.FreqMap[freq]
	if cache == nil {
		cache = NewLRUCache()
		c.FreqMap[freq] = cache
	}
	return cache
}
