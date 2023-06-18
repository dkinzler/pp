package main

// use a trie to store the words in reverse order
// store the stream as a slice
// whenever a new char comes in do a backward search

type StreamChecker struct {
	trie   *TrieNode
	stream []byte
}

func Constructor(words []string) StreamChecker {
	trie := NewTrieNode()
	for _, w := range words {
		trie.AddWordReverse(w)
	}
	return StreamChecker{trie: trie, stream: nil}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.stream = append(this.stream, letter)
	curr := this.trie
	for i := len(this.stream) - 1; i >= 0; i-- {
		next, ok := curr.next[this.stream[i]]
		if !ok {
			break
		}
		if next.isWord {
			return true
		}
		curr = next
	}
	return false
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */

type TrieNode struct {
	isWord bool
	next   map[byte]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{isWord: false, next: make(map[byte]*TrieNode)}
}

func (t *TrieNode) AddWordReverse(w string) {
	curr := t
	for i := len(w) - 1; i >= 0; i-- {
		c := w[i]
		if _, ok := curr.next[c]; !ok {
			curr.next[c] = NewTrieNode()
		}
		curr = curr.next[c]
	}
	curr.isWord = true
}
