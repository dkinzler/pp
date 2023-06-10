package main

import "sort"

func maximumScore(scores []int, edges [][]int) int {
	// path of length 4 with maximum score
	// there are 3 edges, consider iterating over each edge
	// and choosing it as the middle
	// what we need is for both endpoints the next edge with highest weight
	// so what we should do is create an adjacency list for each node
	// and sort it by value

	n := len(scores)
	adj := make([][]E, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a] = append(adj[a], E{b, scores[b]})
		adj[b] = append(adj[b], E{a, scores[a]})
	}

	for i := 0; i < n; i++ {
		sort.Slice(adj[i], func(a, b int) bool {
			return adj[i][a].score > adj[i][b].score
		})
	}

	result := -1
	for _, edge := range edges {
		a, b := edge[0], edge[1]

		c, cs := findBiggest(b, -1, adj[a])
		if c != -1 {
			d, ds := findBiggest(a, c, adj[b])
			if d != -1 {
				s := cs + ds + scores[a] + scores[b]
				if s > result {
					result = s
				}
			}
		}

		c, cs = findBiggest(a, -1, adj[b])
		if c != -1 {
			d, ds := findBiggest(b, c, adj[a])
			if d != -1 {
				s := cs + ds + scores[a] + scores[b]
				if s > result {
					result = s
				}
			}
		}
	}

	return result
}

func findBiggest(e1, e2 int, adj []E) (int, int) {
	for _, e := range adj {
		if e.node != e1 && e.node != e2 {
			return e.node, e.score
		}
	}
	return -1, -1
}

type E struct {
	node  int
	score int
}
