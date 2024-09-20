package trie

type Trie struct {
	Val  string
	Next map[byte]*Trie
}

func (t *Trie) Insert(word string) {
	currTrie := t

	for i := 0; i < len(word); i++ {
		char := word[i]

		if _, ok := currTrie.Next[char]; !ok {
			currTrie.Next[char] = &Trie{Next: make(map[byte]*Trie)}
		}

		currTrie = currTrie.Next[char]
	}

	currTrie.Val = word
}

func (t *Trie) Search(word string) bool {
	currTrie := t

	for i := 0; i < len(word); i++ {
		char := word[i]

		if _, ok := currTrie.Next[char]; !ok {
			return false
		}

		currTrie = currTrie.Next[char]
	}

	return currTrie.Val == word
}

func (t *Trie) FindPrefix(prefix string) bool {
	currTrie := t

	for i := 0; i < len(prefix); i++ {
		char := prefix[i]

		if _, ok := currTrie.Next[char]; !ok {
			return false
		}

		currTrie = currTrie.Next[char]
	}

	return true
}
