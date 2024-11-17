package trie

type Trie struct {
	IsWord   bool
	Children map[rune]*Trie
}

func NewTrie() *Trie {
	return &Trie{Children: make(map[rune]*Trie)}
}

func (t *Trie) Insert(word string) {
	wordRunes := []rune(word)
	n := len(wordRunes)

	var dfs func(node *Trie, idx int)
	dfs = func(node *Trie, idx int) {
		if idx == n {
			node.IsWord = true
			return
		}

		char := wordRunes[idx]
		if _, ok := node.Children[char]; !ok {
			node.Children[char] = NewTrie()
		}

		dfs(node.Children[char], idx+1)
	}

	dfs(t, 0)
}

func (t *Trie) Search(word string) bool {
	var dfs func(trie *Trie, idx int) bool
	dfs = func(trie *Trie, idx int) bool {
		if idx == len(word) {
			return trie.Val == word
		}

		char := word[idx]
		if _, ok := trie.Children[char]; !ok {
			return false
		}

		return dfs(trie.Children[char], idx+1)
	}

	return dfs(t, 0)
}

func (t *Trie) FindPrefix(prefix string) bool {
	var dfs func(trie *Trie, idx int) bool
	dfs = func(trie *Trie, idx int) bool {
		if idx == len(prefix) {
			return true
		}

		char := prefix[idx]
		if _, ok := trie.Children[char]; !ok {
			return false
		}

		return dfs(trie.Children[char], idx+1)
	}

	return dfs(t, 0)
}
