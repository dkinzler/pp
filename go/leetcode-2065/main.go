package main

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	// we can just do brute-force DFS = backtracking
	// maxTime <= 100 and each edge has time >= 10, so the search tree
	// has depth <= 10
	// since each node has <= 4 incident edges there are at most
	// 4^10 = 2^20 ~ 1M nodes in the search tree
	n := len(values)

	adj := make([][]Edge, n)
	for _, edge := range edges {
		a, b, t := edge[0], edge[1], edge[2]
		adj[a] = append(adj[a], Edge{b, t})
		adj[b] = append(adj[b], Edge{a, t})
	}

	used := make([]bool, n)
	used[0] = true
	return rec(0, 0, values[0], adj, values, used, maxTime)
}

func rec(node, time, quality int, adj [][]Edge, values []int, used []bool, maxTime int) int {
	result := 0
	if node == 0 {
		result = quality
	}
	for _, nb := range adj[node] {
		if time+nb.t <= maxTime {
			nq := quality
			wasUsed := used[nb.x]
			if !wasUsed {
				nq += values[nb.x]
				used[nb.x] = true
			}
			result = max(result, rec(nb.x, time+nb.t, nq, adj, values, used, maxTime))
			if !wasUsed {
				used[nb.x] = false
			}
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Edge struct {
	x int
	t int
}
