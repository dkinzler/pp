package main

import (
	"fmt"
	"sort"
)

func main() {
	mat := [][]int{{3, 1}, {3, 4}}
	fmt.Println(maxIncreasingCells(mat))
	mat = [][]int{{1, 1}, {1, 1}}
	fmt.Println(maxIncreasingCells(mat))
	mat = [][]int{{3, 1, 6}, {-9, 5, 7}}
	fmt.Println(maxIncreasingCells(mat))
}

func maxIncreasingCells(mat [][]int) int {
	// find length of longest increasing sequence
	// that can start anywhere in the grid, from a cell we can jump
	// to any other cell in the same row or column
	// sort the values, can then move from highest to lowest
	// for each row and column keep track of the currently lowest values
	m := len(mat)
	n := len(mat[0])

	x := make([]E, 0, m*n)
	for r, row := range mat {
		for c, cell := range row {
			x = append(x, E{r, c, cell})
		}
	}
	sort.Slice(x, func(i, j int) bool {
		return x[i].v > x[j].v
	})

	// for a given row or column which column/row should we jump to
	// the one with the highest value
	// but we need to keep track of the two highest, since one might have the same value
	// as the currently considered cell
	rowMax := make([][]M, m)
	colMax := make([][]M, n)
	result := 0
	for _, e := range x {
		v := 1
		for _, f := range rowMax[e.r] {
			if e.v != f.v {
				v = f.l + 1
				break
			}
		}
		for _, f := range colMax[e.c] {
			if e.v != f.v {
				if f.l+1 > v {
					v = f.l + 1
				}
				break
			}
		}
		if v > result {
			result = v
		}
		rm := rowMax[e.r]
		if len(rm) == 0 {
			rowMax[e.r] = append(rowMax[e.r], M{e.v, v})
		} else if len(rm) == 1 {
			if rm[0].v == e.v {
				if rm[0].l < v {
					rm[0] = M{e.v, v}
				}
			} else {
				rowMax[e.r] = append(rowMax[e.r], M{e.v, v})
				rm = rowMax[e.r]
				if rm[0].l < rm[1].l {
					rm[0], rm[1] = rm[1], rm[0]
				}
			}
		} else {
			if rm[0].l < v {
				if rm[0].v != e.v {
					rm[1] = rm[0]
				}
				rm[0] = M{e.v, v}
			} else if rm[0].v != e.v && rm[1].l < v {
				rm[1] = M{e.v, v}
			}
		}
		cm := colMax[e.c]
		if len(cm) == 0 {
			colMax[e.c] = append(colMax[e.c], M{e.v, v})
		} else if len(cm) == 1 {
			if cm[0].v == e.v {
				if cm[0].l < v {
					cm[0] = M{e.v, v}
				}
			} else {
				colMax[e.c] = append(colMax[e.c], M{e.v, v})
				cm = colMax[e.c]
				if cm[0].l < cm[1].l {
					cm[0], cm[1] = cm[1], cm[0]
				}
			}
		} else {
			if cm[0].l < v {
				if cm[0].v != e.v {
					cm[1] = cm[0]
				}
				cm[0] = M{e.v, v}
			} else if cm[0].v != e.v && cm[1].l < v {
				cm[1] = M{e.v, v}
			}
		}
	}
	return result
}

type E struct {
	r int
	c int
	v int
}

type M struct {
	v int
	l int
}
