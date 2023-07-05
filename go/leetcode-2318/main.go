package main

import "fmt"

func main() {
	fmt.Println(distinctSequences(4))
	fmt.Println(distinctSequences(2))
}

func distinctSequences(n int) int {
	// compute d(i, a, b) = number of sequences of length n
	// that end in a b
	// it is then straightforward
	mod := 1000000007

	if n == 1 {
		return 6
	}

	gcd := [][]bool{
		{false, false, false, false, false, false},
		{false, true, false, true, false, true},
		{false, false, true, false, false, true},
		{false, true, false, true, false, true},
		{false, false, false, false, true, false},
		{false, true, true, true, false, true},
	}
	d1 := make([][]int, 6)
	d2 := make([][]int, 6)
	for i := 0; i < 6; i++ {
		d1[i] = make([]int, 6)
		d2[i] = make([]int, 6)
	}
	for a := 0; a < 6; a++ {
		for b := 0; b < 6; b++ {
			if a != b && !gcd[a][b] {
				d1[a][b] = 1
			}
		}
	}

	for i := 3; i <= n; i++ {
		curr, prev := d2, d1
		if i%2 == 0 {
			curr, prev = d1, d2
		}
		for a := 0; a < 6; a++ {
			for b := 0; b < 6; b++ {
				if a == b || gcd[a][b] {
					continue
				}
				v := 0
				for c := 0; c < 6; c++ {
					if c != b {
						v = (v + prev[c][a]) % mod
					}
				}
				curr[a][b] = v
			}
		}
	}

	result := 0
	curr := d1
	if n%2 == 1 {
		curr = d2
	}

	for a := 0; a < 6; a++ {
		for b := 0; b < 6; b++ {
			result = (result + curr[a][b]) % mod
		}
	}

	return result
}
