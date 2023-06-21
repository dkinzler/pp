package main

import (
	"fmt"
	"pp/ds/kvheap"
)

func main() {
	n := 5
	source := 0
	destination := 1
	target := 5
	edges := [][]int{{4, 1, -1}, {2, 0, -1}, {0, 3, -1}, {4, 3, -1}}
	fmt.Println(modifiedGraphEdges(n, edges, source, destination, target))
	n = 3
	source = 0
	destination = 2
	target = 6
	edges = [][]int{{0, 1, -1}, {0, 2, 5}}
	fmt.Println(modifiedGraphEdges(n, edges, source, destination, target))
	n = 4
	source = 0
	destination = 2
	target = 6
	edges = [][]int{{1, 0, 4}, {1, 2, 3}, {2, 3, 5}, {0, 3, -1}}
	fmt.Println(modifiedGraphEdges(n, edges, source, destination, target))
}

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	/*
		Some observations:
		- we can always set the weight of an -1 edge to target+1 so it will never be used in a shortest path
		- if there already is a path without -1 edges that has weight < target it is impossible
		- if we add all -1 edges with a weight of 1 and shortest dist is > target then it is impossible

		Idea:
		- only have <= 100 nodes, so we can do some iteration
		- build the graph bit by bit
		- at first add all the edges with positive weight
		- check if there is a path
			- if < target -> return impossible
			- if == target -> set every -1 edge to weight target+1
			- if > target -> continue
		- add -1 edges one by one, each with weight 1
			- check if there is a path
				- if no, continue
				- if <= target set weight of curr -1 edge so that it is exactly target
					- before this step there was not path with <= target using another -1 edge
					  otherwise we would have already stopped
				- if > target, continue
		- instead of performing a fresh dijkstra at every step we can just update
	*/

	adj := make([][]Edge, n)
	result := make([][]int, len(edges))
	openEdges := make([]OpenEdge, 0)
	for i, edge := range edges {
		result[i] = edge
		a, b, w := edge[0], edge[1], edge[2]
		adj[a] = append(adj[a], Edge{b, a, w, i})
		adj[b] = append(adj[b], Edge{a, b, w, i})
		if w == -1 {
			openEdges = append(openEdges, OpenEdge{a, len(adj[a]) - 1, b, len(adj[b]) - 1, w, i})
		}
	}

	d := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = -1
	}
	d[source] = 0
	dijkstra(source, d, adj)
	if d[destination] != -1 {
		if d[destination] < target {
			return [][]int{}
		} else if d[destination] == target {
			for _, e := range openEdges {
				result[e.i][2] = target + 1
			}
			return result
		}
	}
	for i, e := range openEdges {
		x, y := e.x, e.y
		adj[x][e.xi].w = 1
		adj[y][e.yi].w = 1
		result[e.i][2] = 1
		if d[x] != -1 {
			if d[y] == -1 {
				d[y] = d[x] + 1
				dijkstra(y, d, adj)
			} else {
				if d[x] < d[y] {
					d[y] = d[x] + 1
					dijkstra(y, d, adj)
				} else {
					d[x] = min(d[x], d[y]+1)
					dijkstra(x, d, adj)
				}
			}
		} else if d[y] != -1 {
			// d[x] == -1
			d[x] = d[y] + 1
			dijkstra(x, d, adj)
		}
		if d[destination] != -1 {
			if d[destination] <= target {
				result[e.i][2] += target - d[destination]
				for _, f := range openEdges[i+1:] {
					result[f.i][2] = target + 1
				}
				return result
			}
		}
	}

	return [][]int{}
}

func dijkstra(startNode int, d []int, adj [][]Edge) {
	q := kvheap.New[int, int](func(i1, i2 int) bool { return i1 < i2 })
	q.Update(startNode, d[startNode])
	for q.Size() > 0 {
		curr, cd, _ := q.Pop()
		for _, e := range adj[curr] {
			if e.w != -1 {
				newD := cd + e.w
				oldD := d[e.x]
				if oldD == -1 || newD < oldD {
					d[e.x] = newD
					q.Update(e.x, newD)
				}
			}
		}
	}
}

type Edge struct {
	x int
	y int
	w int
	i int
}

type OpenEdge struct {
	x  int
	xi int
	y  int
	yi int
	w  int
	i  int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
