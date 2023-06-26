package main

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// build graph
	// bfs to find shortest path from endWord to all other words
	// or at least to beginWord and any nodes that have dist 1 less
	// thus they might be involved on a shortest path
	// do dfs from begin, for each neighbor if it
	// has correct dist to endWord recurse on it

	// beginWord might or might not be in wordList
	// if endWord is not in wordList there is no solution
	n := len(wordList)

	endWordIndex := -1
	beginWordIndex := -1
	for i, word := range wordList {
		if word == endWord {
			endWordIndex = i
		} else if word == beginWord {
			beginWordIndex = i
		}
	}
	if endWordIndex == -1 {
		return [][]string{}
	}
	if beginWordIndex == -1 {
		beginWordIndex = n
		n += 1
		wordList = append(wordList, beginWord)
	}

	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		u := wordList[i]
		for j := i + 1; j < n; j++ {
			v := wordList[j]
			if hasEdge(u, v) {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}

	d := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = -1
	}
	d[endWordIndex] = 0
	q := []int{endWordIndex}
	for len(q) > 0 {
		curr := q[0]
		if curr == beginWordIndex {
			break
		}
		q = q[1:]
		for _, nb := range adj[curr] {
			if d[nb] == -1 {
				d[nb] = d[curr] + 1
				q = append(q, nb)
			}
		}
	}
	if d[beginWordIndex] == -1 {
		return [][]string{}
	}

	return dfs(beginWordIndex, d[beginWordIndex], d, adj, wordList)
}

func dfs(node int, targetDist int, d []int, adj [][]int, wordList []string) [][]string {
	if targetDist == 0 {
		return [][]string{{wordList[node]}}
	}

	result := [][]string{}
	cw := wordList[node]
	for _, nb := range adj[node] {
		if d[nb] == targetDist-1 {
			r := dfs(nb, targetDist-1, d, adj, wordList)
			for _, v := range r {
				z := append([]string{cw}, v...)
				result = append(result, z)
			}
		}
	}
	return result
}

func hasEdge(u, v string) bool {
	diff := 0
	for i := 0; i < len(u); i++ {
		if u[i] != v[i] {
			diff += 1
			if diff > 1 {
				return false
			}
		}
	}
	return true
}
