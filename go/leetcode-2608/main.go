package main

import "fmt"

func main() {
	n := 7
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}, {3, 4}, {4, 5}, {5, 6}, {6, 3}}
	fmt.Println(findShortestCycle(n, edges))
	n = 6
	edges = [][]int{{4, 2}, {5, 1}, {5, 0}, {0, 3}, {5, 2}, {1, 4}, {1, 3}, {3, 4}}
	fmt.Println(findShortestCycle(n, edges))
	n = 5
	edges = [][]int{{2, 4}, {0, 1}, {3, 2}, {4, 0}, {1, 3}}
	fmt.Println(findShortestCycle(n, edges))
}

func findShortestCycle(n int, edges [][]int) int {
	// number of nodes and edges is small, <= 1000
	// so we can just do bfs
	// if we encounter a node that has already been visited, we have a cycle
	// we can then follow from both endpoints up to find the common ancestor
	// the number of edges along the way gives us the length of the cycle

	// the ancestor for each node in the bf traversal
	ancestor := make([]int, n)
	depth := make([]int, n)
	for i := 0; i < n; i++ {
		ancestor[i] = -1
		depth[i] = -1
	}

	adj := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	shortestCycle := -1

	for i := 0; i < n; i++ {
		if depth[i] != -1 {
			continue
		}

		q := []int{i}
		depth[i] = 0

		for len(q) > 0 {
			curr := q[0]
			q = q[1:]
			for _, nb := range adj[curr] {
				if nb == ancestor[curr] {
					continue
				}
				if depth[nb] != -1 {
					// this node was visited before
					// find common ancestor
					cycleLength := findShortestPath(curr, nb, adj) + 1
					if shortestCycle == -1 || cycleLength < shortestCycle {
						shortestCycle = cycleLength
					}
				} else {
					depth[nb] = depth[curr] + 1
					ancestor[nb] = curr
					q = append(q, nb)
				}
			}
		}
	}
	return shortestCycle
}

func findShortestPath(start, end int, adj [][]int) int {
	// find length of shortest path from start to end using bfs
	// an edge between start and end cannot be used
	d := make([]int, len(adj))

	q := []int{start}
	d[start] = 1
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, nb := range adj[curr] {
			if nb == end {
				if curr != start {
					return d[curr]
				} else {
					continue
				}
			}
			if d[nb] == 0 {
				d[nb] = d[curr] + 1
				q = append(q, nb)
			}
		}
	}
	return 0
}
