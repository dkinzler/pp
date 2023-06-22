package main

import "pp/ds/kvheap"

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	d := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = -1
	}

	adj := make([][]E, n)
	for _, meeting := range meetings {
		x, y, t := meeting[0], meeting[1], meeting[2]
		adj[x] = append(adj[x], E{y, t})
		adj[y] = append(adj[y], E{x, t})
	}

	d[0] = 0
	d[firstPerson] = 0
	result := make([]int, 0)
	q := kvheap.New[int, int](func(i, j int) bool {
		return i < j
	})
	q.Update(0, 0)
	q.Update(firstPerson, 0)
	for q.Size() > 0 {
		curr, _, _ := q.Pop()
		result = append(result, curr)
		for _, nb := range adj[curr] {
			if d[curr] <= nb.t {
				tt := d[nb.x]
				if tt == -1 || nb.t < tt {
					d[nb.x] = nb.t
					q.Update(nb.x, nb.t)
				}
			}
		}
	}

	return result
}

type E struct {
	x int
	t int
}
