package main

import "fmt"

func main() {
	edges := [][]int{{1, 2}, {1, 3}, {2, 3}}
	fmt.Println(findRedundantDirectedConnection(edges))
	edges = [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}
	fmt.Println(findRedundantDirectedConnection(edges))
}

func findRedundantDirectedConnection(edges [][]int) []int {
	// can just brute force remove each edge and do dfs
	// if we detect a cycle that edge is not a solution
	// there should be exactly one node with in-degree 0
	// that node is the root
	// if there is not, removing that one edge should create one
	n := len(edges)

	adj := make([][]int, n)
	inDegree := make([]int, n)

	for _, edge := range edges {
		a, b := edge[0]-1, edge[1]-1
		adj[a] = append(adj[a], b)
		inDegree[b] += 1
	}

	root := -1
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			root = i
		}
	}

	var result []int
	for _, edge := range edges {
		a, b := edge[0]-1, edge[1]-1
		r := root
		if r == -1 {
			if inDegree[b] == 1 {
				r = b
			} else {
				continue
			}
		}
		visited := make([]bool, n)
		count, ok := dfs(r, a, b, adj, visited)
		if ok && count == n {
			result = edge
		}
	}

	return result
}

func dfs(node int, r1, r2 int, adj [][]int, visited []bool) (int, bool) {
	if visited[node] {
		return 0, false
	}
	visited[node] = true
	if len(adj[node]) == 0 || (node == r1 && len(adj[node]) == 1) {
		return 1, true
	}

	count := 1
	for _, nb := range adj[node] {
		if node == r1 && nb == r2 {
			continue
		}
		c, ok := dfs(nb, r1, r2, adj, visited)
		if !ok {
			return 0, false
		}
		count += c
	}

	return count, true
}
