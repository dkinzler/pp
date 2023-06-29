package main

import (
	"pp/ds/kvheap"
)

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// Dijkstra but we visit nodes with higher dist=probability first
	adj := make([][]Edge, n)
	for i, edge := range edges {
		x, y := edge[0], edge[1]
		adj[x] = append(adj[x], Edge{y, succProb[i]})
		adj[y] = append(adj[y], Edge{x, succProb[i]})
	}

	q := kvheap.New[int, float64](func(i1, i2 float64) bool {
		return i1 > i2
	})
	dist := make([]float64, n)
	for i := 0; i < n; i++ {
		dist[i] = 0
	}
	dist[start] = 1
	q.Update(start, 1)

	for q.Size() > 0 {
		node, d, _ := q.Pop()
		for _, e := range adj[node] {
			if d*e.p > dist[e.x] {
				dist[e.x] = d * e.p
				q.Update(e.x, dist[e.x])
			}
		}
	}

	return dist[end]
}

type Edge struct {
	x int
	p float64
}
