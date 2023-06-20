package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 5
	edges := [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}}
	fmt.Println(findCriticalAndPseudoCriticalEdges(n, edges))
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	/*
		in a MST a critical edge is one whose deletion would result in a MST
		with higher total weight, i.e. it is part of every MST
		a pseudo critical edge is one that appears in some MSTs but not all
		i.e. it can be interchanged with another edge of the same weight

		we can sort the edges and build the MST using union-find
		whenever a new edge is added to the MST it is a candidate for a critical edge
		if both endpoints u, v of an edge lie in the same component of the current MST
		then we find all edges on the path between u and v in the current MST
		that have the same weight as the edge (u,v), those edges are no longer critical but
		pseudo-critical
		we can find these using dfs
	*/
	m := len(edges)

	eg := make([]Edge, m)
	for i, e := range edges {
		eg[i] = Edge{e[0], e[1], e[2], i}
	}
	sort.Slice(eg, func(i, j int) bool {
		return eg[i].w < eg[j].w
	})

	mst := make([][]Edge, n)
	uf := make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}

	// 1 = critical, 2 = pseudo-critical
	critical := make([]int, m)

	maxEdge := 0
	for _, edge := range eg {
		a, b := find(edge.x, uf), find(edge.y, uf)
		if a != b {
			union(a, b, uf)
			mst[edge.x] = append(mst[edge.x], edge)
			mst[edge.y] = append(mst[edge.y], edge)
			critical[edge.i] = 1
			maxEdge = max(maxEdge, edge.w)
		} else if maxEdge == edge.w {
			// find edges on path from x to y with weight w
			// and mark those as pseudo edges
			// also mark this edge as pseudo if there are any
			r, _ := dfs(edge.x, -1, mst, edge.y, edge.w)
			if len(r) > 0 {
				for _, edgeIndex := range r {
					critical[edgeIndex] = 2
				}
				critical[edge.i] = 2
			}
		}
	}

	result := make([][]int, 2)
	for i, c := range critical {
		if c == 1 {
			result[0] = append(result[0], i)
		} else if c == 2 {
			result[1] = append(result[1], i)
		}
	}
	return result
}

func dfs(node, parent int, adj [][]Edge, target int, weight int) ([]int, bool) {
	if node == target {
		return nil, true
	}

	var result []int
	found := false
	for _, edge := range adj[node] {
		var other int
		if edge.x == node {
			other = edge.y
		} else {
			other = edge.x
		}
		if other != parent {
			r, ok := dfs(other, node, adj, target, weight)
			if ok {
				found = true
				result = r
				if edge.w == weight {
					result = append(result, edge.i)
				}
				break
			}
		}
	}
	return result, found
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Edge struct {
	x int
	y int
	w int
	i int
}

func find(x int, uf []int) int {
	if uf[x] == x {
		return x
	}
	uf[x] = find(uf[x], uf)
	return uf[x]
}

func union(x, y int, uf []int) {
	uf[x] = y
}
