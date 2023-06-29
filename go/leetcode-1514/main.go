package main

import (
	"fmt"
	"pp/ds/kvheap"
)

func main() {
	n := 5
	edges := [][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}}
	fmt.Println(countRestrictedPaths(n, edges))
}

func countRestrictedPaths(n int, edges [][]int) int {
	// count number of restricted paths from node 0 to n-1
	// in weighted connected graph
	// restricted path is one where shortest distance to node n-1
	// gets strictly smaller with every edge we take
	// solution: first do Dijkstra from node n-1 and then dfs
	adj := make([][]Edge, n)
	for _, edge := range edges {
		x, y, w := edge[0]-1, edge[1]-1, edge[2]
		adj[x] = append(adj[x], Edge{y, w})
		adj[y] = append(adj[y], Edge{x, w})
	}

	dist := dijkstra(n, adj)

	return dfs(0, 1000000007, adj, dist, map[int]int{})
}

func dfs(node, mod int, adj [][]Edge, dist []int, mem map[int]int) int {
	if node == len(adj)-1 {
		return 1
	}
	if v, ok := mem[node]; ok {
		return v
	}

	result := 0
	for _, e := range adj[node] {
		if dist[node] > dist[e.x] {
			result = (result + dfs(e.x, mod, adj, dist, mem)) % mod
		}
	}
	mem[node] = result
	return result
}

func dijkstra(n int, adj [][]Edge) []int {
	q := kvheap.New[int, int](func(i1, i2 int) bool {
		return i1 < i2
	})
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = 1 << 60
	}
	dist[n-1] = 0
	q.Update(n-1, 0)

	for q.Size() > 0 {
		node, d, _ := q.Pop()
		for _, e := range adj[node] {
			if d+e.w < dist[e.x] {
				dist[e.x] = d + e.w
				q.Update(e.x, dist[e.x])
			}
		}
	}

	return dist
}

type Edge struct {
	x int
	w int
}
