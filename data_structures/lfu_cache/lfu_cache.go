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

	lru := c.getOrCreateLRU(newFreq)
	lru.insert(node)
}

// Inserts new lru in FreqMap if not created yet, or return if created
func (c *LFUCache) getOrCreateLRU(freq int) *LRUCache {
	lru, ok := c.FreqMap[freq]
	if !ok { // lru not created yet
		lru = NewLRUCache()
		c.FreqMap[freq] = lru
		return lru
	}

	return lru
}

// // create
// 	push:
// 		accept key and value
// 		generate and insert Node into elems
// 		insert node into lru:
// 			check if lru of freq present and then insert
// 	put:
// func (c *LFUCache) Put(key int, value int) {

// }
