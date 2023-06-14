package main

import "fmt"

func main() {
	words := []string{"abc", "ab", "bc", "b"}
	fmt.Println(sumPrefixScores(words))
}

func sumPrefixScores(words []string) []int {
	// answer[i] = sum for every prefix x of words[i] the number of elements j
	// such that x is a prefix of words[j]
	// use a Trie where each node stores the number of words
	n := len(words)
	result := make([]int, n)
	trie := NewTrieNode()
	for _, w := range words {
		trie.AddWord(w)
	}

	for i, w := range words {
		sum := 0
		curr := trie
		for j := 0; j < len(w); j++ {
			curr = curr.next[w[j]]
			sum += curr.count
		}
		result[i] = sum
	}

	return result
}

type TrieNode struct {
	count int
	next  map[byte]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		count: 0,
		next:  make(map[byte]*TrieNode),
	}
}

func (t *TrieNode) AddWord(word string) {
	t.count += 1
	curr := t
	for i := 0; i < len(word); i++ {
		if _, ok := curr.next[word[i]]; !ok {
			curr.next[word[i]] = NewTrieNode()
		}
		curr = curr.next[word[i]]
		curr.count += 1
	}
}
