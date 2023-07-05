package main

func latestDayToCross(row, col int, cells [][]int) int {
	// just do binary search over the length of cells
	// build a grid where grid[r][c] is time that cells becomes flooded
	grid := make([][]int, row)
	for i := 0; i < row; i++ {
		grid[i] = make([]int, col)
	}
	for i, cell := range cells {
		grid[cell[0]-1][cell[1]-1] = i + 1
	}

	low := 1
	high := row * col
	result := 0
	for low <= high {
		mid := (low + high) / 2
		if isPossible(grid, mid) {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}

func isPossible(grid [][]int, t int) bool {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for r := 0; r < m; r++ {
		visited[r] = make([]bool, n)
	}
	q := make([]Cell, 0)
	for c := 0; c < n; c++ {
		if grid[0][c] > t {
			q = append(q, Cell{0, c})
			visited[0][c] = true
		}
	}
	for len(q) > 0 {
		curr := q[0]
		r, c := curr.r, curr.c
		if r == m-1 {
			return true
		}
		q = q[1:]
		nr, nc := r+1, c
		if isValid(nr, nc, m, n, grid, visited, t) {
			q = append(q, Cell{nr, nc})
			visited[nr][nc] = true
		}
		nr, nc = r-1, c
		if isValid(nr, nc, m, n, grid, visited, t) {
			q = append(q, Cell{nr, nc})
			visited[nr][nc] = true
		}
		nr, nc = r, c+1
		if isValid(nr, nc, m, n, grid, visited, t) {
			q = append(q, Cell{nr, nc})
			visited[nr][nc] = true
		}
		nr, nc = r, c-1
		if isValid(nr, nc, m, n, grid, visited, t) {
			q = append(q, Cell{nr, nc})
			visited[nr][nc] = true
		}
	}

	return false
}

func isValid(r, c, m, n int, grid [][]int, visited [][]bool, t int) bool {
	if r < 0 || c < 0 || r >= m || c >= n {
		return false
	}
	if visited[r][c] {
		return false
	}
	if grid[r][c] <= t {
		return false
	}
	return true
}

type Cell struct {
	r int
	c int
}
