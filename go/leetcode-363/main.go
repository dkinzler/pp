package main

func maxSumSubmatrix(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])

	os := make([][]int, m)
	for r := 0; r < m; r++ {
		os[r] = make([]int, n)
	}
	os[0][0] = matrix[0][0]
	for r := 1; r < m; r++ {
		os[r][0] = matrix[r][0] + os[r-1][0]
	}
	for c := 1; c < n; c++ {
		os[0][c] = matrix[0][c] + os[0][c-1]
	}
	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			os[r][c] = matrix[r][c] + os[r-1][c] + os[r][c-1] - os[r-1][c-1]
		}
	}

	result := 0
	resultSet := false
	for r1 := 0; r1 < m; r1++ {
		for c1 := 0; c1 < n; c1++ {
			for r2 := r1; r2 < m; r2++ {
				for c2 := c1; c2 < n; c2++ {
					v := os[r2][c2]
					if r1 > 0 {
						v -= os[r1-1][c2]
					}
					if c1 > 0 {
						v -= os[r2][c1-1]
					}
					if r1 > 0 && c1 > 0 {
						v += os[r1-1][c1-1]
					}
					if v <= k {
						if !resultSet {
							result = v
							resultSet = true
						} else if v > result {
							result = v
						}
					}
				}
			}
		}
	}
	return result
}
