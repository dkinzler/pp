package main

import "fmt"

func main() {
	s := "aab"
	fmt.Println(minCut(s))
	s = "a"
	fmt.Println(minCut(s))
	s = "ab"
	fmt.Println(minCut(s))
	s = "abbaabb"
	fmt.Println(minCut(s))
}

func minCut(s string) int {
	// precompute d(i, j) = true if s[i...j] is palindromic
	// by iterating over length of [i, j]
	// then do dp again to find c(i) = min cuts needed for palindrome partitioning of s[0...i]
	// iterate over j and if s[j...i] is a palindrome use e[0...j-1]
	n := len(s)
	d := make([][]bool, n)
	for i := 0; i < n; i++ {
		d[i] = make([]bool, n)
		d[i][i] = true
	}
	for k := 2; k <= n; k++ {
		for i := 0; i+k-1 < n; i++ {
			j := i + k - 1
			if s[i] == s[j] {
				if i+1 >= j-1 || d[i+1][j-1] {
					d[i][j] = true
				}
			}
		}
	}

	e := make([]int, n)
	for i := 1; i < n; i++ {
		min := -1
		for j := i; j >= 0; j-- {
			if d[j][i] {
				v := 0
				if j-1 >= 0 {
					v += 1 + e[j-1]
				}
				if min == -1 || v < min {
					min = v
				}
			}
		}
		e[i] = min
	}
	return e[n-1]
}
