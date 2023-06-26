package main

import (
	"fmt"
	"pp/ds/kvheap"
)

func main() {
	heightMap := [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}
	fmt.Println(trapRainWater(heightMap))
	heightMap = [][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}}
	fmt.Println(trapRainWater(heightMap))
}

func trapRainWater(heightMap [][]int) int {
	// compute for each cell (r, c) the smallest edge cell that is reachable from there
	// where an edge cell is a cell that can not hold water
	// initially that is all the outer cells, we keep them in a heap
	// suppose we visit a cell (r, c) with height h
	// and a neighbor with height z, if z >= h then the min edge we can reach from that cell becomes z
	// if z < h then the min edge we can reach becomes z
	// we add these to the heap and continue until we have those heights determined for each cell
	// when we have done this, we can compute the water trapped

	m := len(heightMap)
	n := len(heightMap[0])

	mm := make([][]int, m)
	for r := 0; r < m; r++ {
		mm[r] = make([]int, n)
		for c := 0; c < n; c++ {
			mm[r][c] = -1
		}
	}

	edgeBfs(heightMap, mm, m, n)
	result := 0
	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			h, z := heightMap[r][c], mm[r][c]
			if z > h {
				result += z - h
			}
		}
	}
	return result
}

func edgeBfs(heightMap, mm [][]int, m, n int) {
	q := kvheap.New[Cell, int](func(i1, i2 int) bool {
		return i1 < i2
	})

	for c := 0; c < n; c++ {
		q.Update(Cell{0, c}, heightMap[0][c])
		mm[0][c] = heightMap[0][c]
		if m > 1 {
			q.Update(Cell{m - 1, c}, heightMap[m-1][c])
			mm[m-1][c] = heightMap[m-1][c]
		}
	}
	for r := 1; r < m-1; r++ {
		q.Update(Cell{r, 0}, heightMap[r][0])
		mm[r][0] = heightMap[r][0]
		if n > 1 {
			q.Update(Cell{r, n - 1}, heightMap[r][n-1])
			mm[r][n-1] = heightMap[r][n-1]
		}
	}
	for q.Size() > 0 {
		curr, _, _ := q.Pop()
		r, c := curr.r, curr.c
		h := mm[r][c]
		for _, nb := range [][]int{{r + 1, c}, {r - 1, c}, {r, c + 1}, {r, c - 1}} {
			nr, nc := nb[0], nb[1]
			if !isValid(nr, nc, m, n) {
				continue
			}
			z := mm[nr][nc]
			if z == -1 {
				if heightMap[nr][nc] >= h {
					mm[nr][nc] = heightMap[nr][nc]
					q.Update(Cell{nr, nc}, heightMap[nr][nc])
				} else {
					mm[nr][nc] = h
					q.Update(Cell{nr, nc}, h)
				}
			}
		}
	}
}

func isValid(r, c, m, n int) bool {
	return r >= 0 && c >= 0 && r < m && c < n
}

type Cell struct {
	r int
	c int
}
