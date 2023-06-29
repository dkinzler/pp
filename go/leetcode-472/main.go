package main

import "fmt"

func main() {
	words := []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}
	fmt.Println(findAllConcatenatedWordsInADict(words))
	words = []string{"cat", "dog", "catdog"}
	fmt.Println(findAllConcatenatedWordsInADict(words))
}

func findAllConcatenatedWordsInADict(words []string) []string {
	// find all the words that can be formed by concatenating 2+ shorter words
	// use a trie and dp
	trie := NewTrieNode()
	wordMap := make(map[string]bool)
	for _, w := range words {
		trie.AddWord(w)
		wordMap[w] = true
	}

	result := []string{}
	mem := make(map[string]bool)
	for _, w := range words {
		if dp(w, trie, wordMap, mem) {
			result = append(result, w)
		}
	}
	return result
}

// returns true if word is concatenated
func dp(w string, trie *TrieNode, wordMap map[string]bool, mem map[string]bool) bool {
	if len(w) == 1 {
		return false
	}
	if z, ok := mem[w]; ok {
		return z
	}

	result := false
	curr := trie
	for i := 0; i < len(w)-1; i++ {
		curr = curr.Next(w[i])
		if curr == nil {
			result = false
			break
		}
		if curr.isWord {
			v := w[i+1:]
			if _, ok := wordMap[v]; ok {
				result = true
				break
			} else {
				if dp(v, trie, wordMap, mem) {
					result = true
					break
				}
			}
		}
	}
	mem[w] = result
	return result
}

type TrieNode struct {
	next   map[byte]*TrieNode
	isWord bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		next:   make(map[byte]*TrieNode),
		isWord: false,
	}
}

func (t *TrieNode) AddWord(w string) {
	curr := t
	for i := 0; i < len(w); i++ {
		c := w[i]
		if _, ok := curr.next[c]; !ok {
			curr.next[c] = NewTrieNode()
		}
		curr = curr.next[c]
	}
	curr.isWord = true
}

func (t *TrieNode) Next(c byte) *TrieNode {
	return t.next[c]
}
