package main

func rootCount(edges [][]int, guesses [][]int, k int) int {
	// first do a dfs with node 0 and count the number of correct guesses
	// the key observation is that when node x is the root and has a neighbor y
	// if we change the root to y, at most the edge (x, y) becomes or ceases to be a correct guess
	// so we can again do dfs and count the number of nodes
	n := len(edges) + 1
	adj := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	guessMap := make(map[E]bool)
	for _, guess := range guesses {
		// a is parent of b
		a, b := guess[0], guess[1]
		guessMap[E{a, b}] = true
	}

	initialCorrect := dfs(0, -1, adj, guessMap)
	return dfs2(0, -1, initialCorrect, k, adj, guessMap)
}

func dfs(node, parent int, adj [][]int, guesses map[E]bool) int {
	correct := 0
	if _, ok := guesses[E{parent, node}]; ok {
		correct += 1
	}
	for _, nb := range adj[node] {
		if nb != parent {
			correct += dfs(nb, node, adj, guesses)
		}
	}
	return correct
}

func dfs2(node, parent, correct, k int, adj [][]int, guesses map[E]bool) int {
	result := 0

	if _, ok := guesses[E{parent, node}]; ok {
		correct -= 1
	}
	if _, ok := guesses[E{node, parent}]; ok {
		correct += 1
	}
	if correct >= k {
		result += 1
	}

	for _, nb := range adj[node] {
		if nb != parent {
			result += dfs2(nb, node, correct, k, adj, guesses)
		}
	}

	return result
}

type E struct {
	a int
	b int
}
