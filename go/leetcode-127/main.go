package main

import "fmt"

func main() {
	beginWord, endWord := "hit", "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
	beginWord, endWord = "hit", "cog"
	wordList = []string{"hot", "dot", "dog", "lot", "log"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// isn't this just a bfs among words?
	// maybe we can optimize how to create neighbors and co a bit?
	n := len(wordList)
	si, ei := -1, -1
	for i, v := range wordList {
		if v == endWord {
			ei = i
		}
		if v == beginWord {
			si = i
		}
	}
	// begin and end word are not necessarily in wordList
	if ei == -1 {
		return 0
	}
	if si == -1 {
		si = n
		n += 1
	}

	adj := make([][]int, n)
	for i, v := range wordList {
		for j := i + 1; j < len(wordList); j++ {
			if isNeighbor(v, wordList[j]) {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}
	if si >= len(wordList) {
		for i, v := range wordList {
			if isNeighbor(v, beginWord) {
				adj[i] = append(adj[i], si)
				adj[si] = append(adj[si], i)
			}
		}
	}

	visited := make([]bool, n)
	q := []E{{si, 0}}
	visited[si] = true

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, nb := range adj[curr.i] {
			if !visited[nb] {
				d := curr.d + 1
				if nb == ei {
					return d + 1
				}
				visited[nb] = true
				q = append(q, E{nb, d})
			}
		}
	}

	return 0
}

func isNeighbor(a, b string) bool {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff += 1
		}
	}
	return diff == 1
}

type E struct {
	i int
	d int
}
