package main

import "fmt"

func main() {
	isInfected := [][]int{{0, 1, 0, 0, 0, 0, 0, 1}, {0, 1, 0, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 0, 0, 0}}
	fmt.Println(containVirus(isInfected))
}

func containVirus(isInfected [][]int) int {
	m := len(isInfected)
	n := len(isInfected[0])

	d := make([][]int, m)
	for i := 0; i < m; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			d[i][j] = isInfected[i][j]
		}
	}

	// quarantined cells will be marked with 2

	usedWalls := 0
	for {
		// find components
		comps := []E{}
		visited := map[E]bool{}
		for r := 0; r < m; r++ {
			for c := 0; c < n; c++ {
				if d[r][c] == 1 && !visited[E{r, c}] {
					comps = append(comps, E{r, c})
					findComponent(r, c, d, m, n, visited)
				}
			}
		}
		if len(comps) == 0 {
			break
		}

		// compute threatened cells
		s := []map[E]bool{}
		visited = map[E]bool{}
		max := -1
		maxIndex := -1
		for i, pos := range comps {
			th := map[E]bool{}
			computeThreatened(pos.r, pos.c, d, m, n, th, visited)
			s = append(s, th)
			if len(th) > max {
				max = len(th)
				maxIndex = i
			}
		}

		if max == 0 {
			break
		}

		// quarantine maxIndex
		visited = map[E]bool{}
		usedWalls += computeWalls(comps[maxIndex].r, comps[maxIndex].c, d, m, n, visited)

		// spread infection in other components
		for i, v := range s {
			if i != maxIndex {
				for pos := range v {
					d[pos.r][pos.c] = 1
				}
			}
		}
	}
	return usedWalls
}

func findComponent(r, c int, d [][]int, m, n int, visited map[E]bool) {
	if r < 0 || c < 0 || r >= m || c >= n {
		return
	}
	if d[r][c] == 0 || d[r][c] == 2 || visited[E{r, c}] {
		return
	}
	visited[E{r, c}] = true
	findComponent(r+1, c, d, m, n, visited)
	findComponent(r-1, c, d, m, n, visited)
	findComponent(r, c+1, d, m, n, visited)
	findComponent(r, c-1, d, m, n, visited)
}

type E struct {
	r int
	c int
}

func computeThreatened(r, c int, d [][]int, m, n int, threatened, visited map[E]bool) {
	if r < 0 || c < 0 || r >= m || c >= n {
		return
	}
	if d[r][c] == 2 || visited[E{r, c}] {
		return
	}
	if d[r][c] == 0 {
		threatened[E{r, c}] = true
		return
	}
	visited[E{r, c}] = true
	computeThreatened(r+1, c, d, m, n, threatened, visited)
	computeThreatened(r-1, c, d, m, n, threatened, visited)
	computeThreatened(r, c+1, d, m, n, threatened, visited)
	computeThreatened(r, c-1, d, m, n, threatened, visited)
}

func computeWalls(r, c int, d [][]int, m, n int, visited map[E]bool) int {
	if r < 0 || c < 0 || r >= m || c >= n {
		return 0
	}
	if d[r][c] == 0 {
		return 1
	}
	if d[r][c] == 2 || visited[E{r, c}] {
		return 0
	}
	visited[E{r, c}] = true
	d[r][c] = 2
	result := 0
	result += computeWalls(r+1, c, d, m, n, visited)
	result += computeWalls(r-1, c, d, m, n, visited)
	result += computeWalls(r, c+1, d, m, n, visited)
	result += computeWalls(r, c-1, d, m, n, visited)
	return result
}
