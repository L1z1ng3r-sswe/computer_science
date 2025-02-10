package main

type MyHashMap struct {
	Buckets []*Bucket
	Len     int // Number of elements
	BLen    int // Number of buckets
	BCap    int // Bucket capacity
}

type Bucket struct {
	NextBucket *Bucket
	Mem        []*Cell
}

type Cell struct {
	Key int
	Val int
}

func Constructor() MyHashMap {
	ht := MyHashMap{
		Buckets: make([]*Bucket, 16), // Default initial bucket size
		BLen:    16,
		BCap:    4,
	}

	for i := range ht.Buckets {
		ht.Buckets[i] = ht.newBucket()
	}

	return ht
}

func (ht *MyHashMap) Get(key int) int {
	bucketIdx, bucketChainIdx, bucketKeyIdx, ok := ht.search(key)
	if !ok {
		return -1
	}

	bucket := ht.Buckets[bucketIdx]
	for i := 0; i < bucketChainIdx; i++ {
		bucket = bucket.NextBucket
	}

	return bucket.Mem[bucketKeyIdx].Val
}

func (ht *MyHashMap) Remove(key int) {
	bucketIdx, bucketChainIdx, bucketKeyIdx, ok := ht.search(key)
	if !ok {
		return
	}

	bucket := ht.Buckets[bucketIdx]
	for i := 0; i < bucketChainIdx; i++ {
		bucket = bucket.NextBucket
	}

	bucket.Mem[bucketKeyIdx] = nil
	ht.Len--
}

func (ht *MyHashMap) search(key int) (int, int, int, bool) {
	bucketIdx := ht.hashFunction(key)
	bucket := ht.Buckets[bucketIdx]
	bucketChainIdx := 0

	for bucket != nil {
		for i, cell := range bucket.Mem {
			if cell != nil && cell.Key == key {
				return bucketIdx, bucketChainIdx, i, true
			}
		}
		bucketChainIdx++
		bucket = bucket.NextBucket
	}
	return -1, -1, -1, false
}

func (ht *MyHashMap) Put(key int, val int) {
	if float64(ht.Len)/float64(ht.BLen) > 6.5 {
		ht.grow()
	}

	bucketIdx := ht.hashFunction(key)
	bucket := ht.Buckets[bucketIdx]

	for {
		for i := 0; i < ht.BCap; i++ {
			if bucket.Mem[i].Key == key {
				bucket.Mem[i] = &Cell{Key: key, Val: val}
				return
			}
		}

		if bucket.NextBucket == nil {
			bucket.NextBucket = ht.newBucket()
		}
		bucket = bucket.NextBucket
	}
}

func (ht *MyHashMap) grow() {
	newSize := ht.BLen * 2
	newBuckets := make([]*Bucket, newSize)

	for i := range newBuckets {
		newBuckets[i] = ht.newBucket()
	}

	for _, bucket := range ht.Buckets {
		for bucket != nil {
			for _, cell := range bucket.Mem {
				if cell != nil {
					newIdx := cell.Key % newSize
					newBucket := newBuckets[newIdx]

					for {
						for i := 0; i < ht.BCap; i++ {
							if newBucket.Mem[i] == nil {
								newBucket.Mem[i] = cell
								break
							}
						}
						if newBucket.NextBucket == nil {
							newBucket.NextBucket = ht.newBucket()
							break
						}
						newBucket = newBucket.NextBucket
					}
				}
			}
			bucket = bucket.NextBucket
		}
	}

	ht.Buckets = newBuckets
	ht.BLen = newSize
}

func (ht *MyHashMap) hashFunction(key int) int {
	return key % ht.BLen
}

func (ht *MyHashMap) newBucket() *Bucket {
	return &Bucket{
		Mem: make([]*Cell, ht.BCap),
	}
}
