package trie

type Trie struct {
	Val  string
	Next map[byte]*Trie
}

func (t *Trie) Insert(word string) {
	var dfs func(trie *Trie, word string, idx int)
	dfs = func(trie *Trie, word string, idx int) {
		if idx == len(word) {
			trie.Val = word
			return
		}

		if trie.Next == nil {
			trie.Next = make(map[byte]*Trie)
		}

		char := word[idx]

		if _, ok := trie.Next[char]; !ok {
			trie.Next[char] = &Trie{Next: make(map[byte]*Trie)}
		}
		dfs(trie.Next[char], word, idx+1)
	}

	dfs(t, word, 0)
}

func (t *Trie) Search(word string) bool {
	var dfs func(trie *Trie, idx int) bool
	dfs = func(trie *Trie, idx int) bool {
		if idx == len(word) {
			return trie.Val == word
		}

		char := word[idx]
		if _, ok := trie.Next[char]; !ok {
			return false
		}

		return dfs(trie.Next[char], idx+1)
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
		if _, ok := trie.Next[char]; !ok {
			return false
		}

		return dfs(trie.Next[char], idx+1)
	}

	return dfs(t, 0)
}
