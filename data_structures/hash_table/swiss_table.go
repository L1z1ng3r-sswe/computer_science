package main

type Map struct {
	tables []table
}

type table struct {
	prefix     uint8 // N upper bits a table covers
	prefixBits int8  // how many upper bits to check
	groups     []entry
	ctrls      []byte
}

type entry struct {
	key   string
	value int
}

// Insert function updates a value if a key already exists
func (m Map) Insert(key string, val int) {

}

func (m Map) Delete(key string) {

}

// Search function if doesn't find the key the argument returned would be false
func (m Map) Search(key string) (int, bool) {

}

func (m Map) getHash(key string) uint64 {

}
