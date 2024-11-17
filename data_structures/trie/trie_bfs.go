package trie

type Trie struct {
	IsWord bool
	Next   map[rune]*Trie
}

func NewTrie() *Trie {
	return &Trie{Next: make(map[rune]*Trie)}
}

func (t *Trie) Insert(word string) {
	currTrie := t

	for _, char := range word {
		if _, ok := currTrie.Next[char]; !ok {
			currTrie.Next[char] = NewTrie()
		}
		currTrie = currTrie.Next[char]
	}

	currTrie.IsWord = true
}

func (t *Trie) Search(word string) bool {
	currTrie := t

	for _, char := range word {

		if _, ok := currTrie.Next[char]; !ok {
			return false
		}

		currTrie = currTrie.Next[char]
	}

	return currTrie.IsWord
}

func (t *Trie) FindPrefix(prefix string) bool {
	currTrie := t

	for _, char := range prefix {

		if _, ok := currTrie.Next[char]; !ok {
			return false
		}

		currTrie = currTrie.Next[char]
	}

	return true
}
