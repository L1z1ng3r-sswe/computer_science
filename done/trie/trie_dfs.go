type Trie struct {
	Exist bool
	Next  map[byte]*Trie
}

func Constructor() Trie {
	return Trie{Exist: false, Next: make(map[byte]*Trie)}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.Exist = true
		return
	}

	firstByte := word[0]
	nextByte, ok := this.Next[firstByte]
	if !ok {
		nextByte = &Trie{Next: make(map[byte]*Trie)}
		this.Next[firstByte] = nextByte
	}

	nextByte.Insert(word[1:])
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.Exist
	}

	firstByte := word[0]
	nextByte, ok := this.Next[firstByte]
	if !ok {
		return false
	}

	return nextByte.Search(word[1:])
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}

	firstByte := prefix[0]
	nextByte, ok := this.Next[firstByte]
	if !ok {
		return false
	}

	return nextByte.StartsWith(prefix[1:])
}



git add .
git commit -m "feat: 208_trie_data_structure done; approach: dfs"
git push origin main
