package main

import "fmt"

func main() {
	s, t := "rabbbit", "rabbit"
	fmt.Println(numDistinct(s, t))
	s, t = "babgbag", "bag"
	fmt.Println(numDistinct(s, t))
}

// compute the number of distinct subsequences of s equal to t
func numDistinct(s string, t string) int {
	n := len(s)
	m := len(t)

	// this is a classic DP, we solve it bottom-up
	// d(i, j) = number of subsequences of s[0...i] which equal t[0...j]
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, m)
	}

	if s[0] == t[0] {
		d[0][0] = 1
	}
	for i := 1; i < n; i++ {
		v := d[i-1][0]
		if s[i] == t[0] {
			v += 1
		}
		d[i][0] = v
	}

	for j := 1; j < m; j++ {
		for i := j; i < n; i++ {
			v := d[i-1][j]
			if s[i] == t[j] {
				v += d[i-1][j-1]
			}
			d[i][j] = v
		}
	}

	return d[n-1][m-1]
}
