package main

import (
	"fmt"
	"os"
	"pp/ds/kvheap"
	"pp/io"
)

func main() {
	reader, err := io.NewLineReader("testcases.txt")
	defer reader.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	i := 0
	for {
		c, err := readTestCase(reader)
		if err != nil {
			if err != io.ErrorEOF {
				fmt.Println(err)
			}
			return
		}
		result := minimumObstacles(c.grid)
		if result == c.expected {
			fmt.Println("case", i, "ok")
		} else {
			fmt.Printf("case %v incorrect, expected %v got %v\n", i, c.expected, result)
		}
		i += 1
	}
}

type testCase struct {
	grid     [][]int
	expected int
}

func readTestCase(reader *io.LineReader) (testCase, error) {
	var result testCase
	grid, err := reader.Read2DimIntSlice()
	if err != nil {
		return result, err
	}
	expected, err := reader.ReadInt()
	if err != nil {
		return result, err
	}
	result.grid = grid
	result.expected = expected
	return result, nil
}

func minimumObstacles(grid [][]int) int {
	// this is a classic dijkstra
	// where we view the grid as a graph, between empty cells use
	// edge weight 0, to a cell with an obstacle use weight 1
	// length of shortest path from 0,0 to m-1,n-1 is then min number of
	// obstacles to remove
	m := len(grid)
	n := len(grid[0])

	d := make([][]int, m)
	for i := 0; i < m; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			d[i][j] = -1
		}
	}
	d[0][0] = 0
	q := kvheap.New[Pos, int](func(i1, i2 int) bool { return i1 < i2 })
	q.Update(Pos{0, 0}, 0)

	nbs := func(i Pos, j int) (Pos, bool, bool) {
		r, c := i.r, i.c
		if j == 0 {
			r += 1
		} else if j == 1 {
			r -= 1
		} else if j == 2 {
			c += 1
		} else {
			c -= 1
		}
		if r < 0 || c < 0 || r >= m || c >= n {
			return Pos{}, false, false
		}
		ni := Pos{r, c}
		if grid[r][c] == 1 {
			return ni, true, true
		} else {
			return ni, false, true
		}
	}

	for q.Size() > 0 {
		curr, _, _ := q.Pop()
		dist := d[curr.r][curr.c]
		for i := 0; i < 4; i++ {
			nb, isObstacle, ok := nbs(curr, i)
			if ok {
				z := dist
				if isObstacle {
					z += 1
				}
				old := d[nb.r][nb.c]
				if old == -1 || z < old {
					d[nb.r][nb.c] = z
					q.Update(nb, z)
				}
			}
		}
	}

	return d[m-1][n-1]
}

type Pos struct {
	r int
	c int
}
