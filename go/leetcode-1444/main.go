package main

import "fmt"

func main() {
	pizza := []string{"A..", "AAA", "..."}
	k := 3
	fmt.Println(ways(pizza, k))
}

func ways(pizza []string, k int) int {
	// <= 50 rows, cols
	// note that since we always give away the left or upper part
	// the bottom right cell will always be remaining
	// i.e. we can uniquely identify every possible remaining part
	// by the top left cell
	// solve it using DP, could do either top-down or bottom-up
	m := len(pizza)
	n := len(pizza[0])
	count := make([][]int, m)
	for i := 0; i < m; i++ {
		count[i] = make([]int, n)
	}
	if pizza[m-1][n-1] == byte('A') {
		count[m-1][n-1] = 1
	}
	for j := n - 2; j >= 0; j-- {
		count[m-1][j] = count[m-1][j+1]
		if pizza[m-1][j] == byte('A') {
			count[m-1][j] += 1
		}
	}
	for i := m - 2; i >= 0; i-- {
		count[i][n-1] = count[i+1][n-1]
		if pizza[i][n-1] == byte('A') {
			count[i][n-1] += 1
		}
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			v := 0
			if pizza[i][j] == byte('A') {
				v += 1
			}
			v += count[i+1][j] + count[i][j+1] - count[i+1][j+1]
			count[i][j] = v
		}
	}

	mem := make(map[E]int)
	return rec(0, 0, k, count, mem)
}

// returns the number of ways to divide
// the rectangle with top left cell (r,c) into k pieces
func rec(r, c, k int, count [][]int, mem map[E]int) int {
	if k == 1 {
		if count[r][c] > 0 {
			return 1
		}
		return 0
	}

	mod := 1000000007

	if v, ok := mem[E{r, c, k}]; ok {
		return v
	}

	result := 0
	// horizontal cuts
	for h := c + 1; h < len(count[0]); h++ {
		if (count[r][c] - count[r][h]) > 0 {
			result = (result + rec(r, h, k-1, count, mem)) % mod
		}
	}
	// vertical cuts
	for v := r + 1; v < len(count); v++ {
		if (count[r][c] - count[v][c]) > 0 {
			result = (result + rec(v, c, k-1, count, mem)) % mod
		}
	}

	mem[E{r, c, k}] = result
	return result
}

type E struct {
	r int
	c int
	k int
}
