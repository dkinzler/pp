package main

import "fmt"

func main() {
	fmt.Println(countVowelPermutation(1))
	fmt.Println(countVowelPermutation(2))
	fmt.Println(countVowelPermutation(5))
}

func countVowelPermutation(n int) int {
	// compute how many strings of length n are there that only contain vowels a, e, i, o, u
	// where each vowel can only be followed by some (not all) other vowels
	const a, e, i, o, u int = 0, 1, 2, 3, 4
	k := 5
	mod := 1000000007
	// compute d(i, j) = number of strings of length i that end in vowel j
	d := make([][]int, 2)
	d[0] = make([]int, k)
	d[1] = make([]int, k)
	// strings of length 1
	for j := 0; j < k; j++ {
		d[0][j] = 1
	}
	for j := 1; j < n; j++ {
		curr := j % 2
		prev := 1 - curr
		// ends in a, before can be e, i, u
		d[curr][a] = (d[prev][e] + d[prev][i] + d[prev][u]) % mod
		// ends in e, before can be a, i
		d[curr][e] = (d[prev][a] + d[prev][i]) % mod
		// ends in i, before can be e, o
		d[curr][i] = (d[prev][e] + d[prev][o]) % mod
		// ends in o, before can be i
		d[curr][o] = d[prev][i]
		// ends in u, before can be i, o
		d[curr][u] = (d[prev][i] + d[prev][o]) % mod
	}

	result := 0
	last := (n - 1) % 2
	for j := 0; j < k; j++ {
		result = (result + d[last][j]) % mod
	}

	return result
}
