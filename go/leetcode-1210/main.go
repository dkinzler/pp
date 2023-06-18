package main

func minimumMoves(grid [][]int) int {
	n := len(grid)
	dist := make(map[Pos]int)
	q := []Pos{{0, 0, 0, 1}}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		d := dist[curr]
		for _, nb := range curr.Neighbors(grid) {
			if _, ok := dist[nb]; !ok {
				dist[nb] = d + 1
				q = append(q, nb)
				if isGoal(nb, n) {
					return d + 1
				}
			}
		}
	}

	return -1
}

func isGoal(p Pos, n int) bool {
	return p.r1 == n-1 && p.c1 == n-2 && p.r2 == n-1 && p.c2 == n-1
}

type Pos struct {
	r1 int
	c1 int
	r2 int
	c2 int
}

func (p *Pos) Neighbors(grid [][]int) []Pos {
	var result []Pos
	n := len(grid)
	// move right
	if p.c2 < n-1 {
		nr1, nc1 := p.r1, p.c1+1
		nr2, nc2 := p.r2, p.c2+1
		if grid[nr1][nc1] == 0 && grid[nr2][nc2] == 0 {
			result = append(result, Pos{nr1, nc1, nr2, nc2})
		}
	}
	// move down
	if p.r2 < n-1 {
		nr1, nc1 := p.r1+1, p.c1
		nr2, nc2 := p.r2+1, p.c2
		if grid[nr1][nc1] == 0 && grid[nr2][nc2] == 0 {
			result = append(result, Pos{nr1, nc1, nr2, nc2})
		}
	}
	// rotate clockwise
	if p.r1 == p.r2 && p.r1 < n-1 {
		nr1, nc1 := p.r1, p.c1
		nr2, nc2 := p.r2+1, p.c2-1
		hr, hc := p.r2+1, p.c2
		if grid[hr][hc] == 0 && grid[nr2][nc2] == 0 {
			result = append(result, Pos{nr1, nc1, nr2, nc2})
		}
	}
	// rotate clockwise
	if p.c1 == p.c2 && p.c1 < n-1 {
		nr1, nc1 := p.r1, p.c1
		nr2, nc2 := p.r2-1, p.c2+1
		hr, hc := p.r2, p.c2+1
		if grid[hr][hc] == 0 && grid[nr2][nc2] == 0 {
			result = append(result, Pos{nr1, nc1, nr2, nc2})
		}
	}
	return result
}
