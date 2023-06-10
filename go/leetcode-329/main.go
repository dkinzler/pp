package main

import "sort"

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	d := make([][]int, m)
	for i := 0; i < m; i++ {
		d[i] = make([]int, n)
	}

	q := make([]E, 0, m*n)
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			q = append(q, E{r, c, matrix[r][c]})
		}
	}

	sort.Slice(q, func(i, j int) bool {
		return q[i].v < q[j].v
	})

	result := 1
	for _, e := range q {
		r, c, v := e.r, e.c, e.v
		z := 1
		if r > 0 && matrix[r-1][c] < v {
			z = max(z, d[r-1][c]+1)
		}
		if r < m-1 && matrix[r+1][c] < v {
			z = max(z, d[r+1][c]+1)
		}
		if c > 0 && matrix[r][c-1] < v {
			z = max(z, d[r][c-1]+1)
		}
		if c < n-1 && matrix[r][c+1] < v {
			z = max(z, d[r][c+1]+1)
		}
		d[r][c] = z
		if z > result {
			result = z
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

type E struct {
	r int
	c int
	v int
}
