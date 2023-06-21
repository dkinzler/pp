package main

import "fmt"

func main() {
	graph := [][]int{{1, 2, 3}, {0}, {0}, {0}}
	fmt.Println(shortestPathLength(graph))
	graph = [][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}}
	fmt.Println(shortestPathLength(graph))
}

func shortestPathLength(graph [][]int) int {
	// let d(i, j) = length of shortest path to visit all nodes in bitmask i
	// and end the path at node j
	// d(1<<j, j) = 0 for all nodes j
	// we can do bfs starting from those elements
	// from d(i, j) we can reach d(i|(1<<k), k) for all neighbors of node j
	// an element (i, j) will enter the bfs queue at most once
	// and for each element in the q we need to iterate over its neighbors
	// so in total this will take <= 2^12*12*12 <= 5k * 200 = 1M  steps
	n := len(graph)
	nn := 1 << n
	all := nn - 1

	q := make([]E, 0, n)
	d := make([][]int, nn)
	for i := 0; i < nn; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			d[i][j] = -1
		}
	}
	for i := 0; i < n; i++ {
		q = append(q, E{1 << i, i, 0})
		d[1<<i][i] = 0
	}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, nb := range graph[curr.e] {
			s := curr.s | (1 << nb)
			if s == all {
				return curr.d + 1
			}
			v := curr.d + 1
			if d[s][nb] == -1 {
				d[s][nb] = v
				q = append(q, E{s, nb, v})
			}
		}
	}
	return 0
}

type E struct {
	s int
	e int
	d int
}
