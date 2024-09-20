package main

import "fmt"

func main() {
	trie := &Trie{Next: make(map[rune]*Trie)}

	trie.Insert("as")
	trie.Insert("asy")
	trie.Insert("asyl")
	trie.Insert("asylbek")

	fmt.Println(trie.Search(""))
}

type Trie struct {
	Val  string
	Next map[rune]*Trie
}

func (t *Trie) Insert(word string) {
	var currTrie = t

	for _, char := range word {
		if _, ok := currTrie.Next[char]; !ok {
			currTrie.Next[char] = &Trie{Next: make(map[rune]*Trie)}
		}

		currTrie = currTrie.Next[char]
	}

	currTrie.Val = word
}

func (t *Trie) Search(word string) bool {
	var currTrie = t

	for _, char := range word {
		if _, ok := currTrie.Next[char]; !ok {
			return false
		}

		currTrie = currTrie.Next[char]
	}

	return currTrie.Val == word
}
