package main

import "fmt"

func main() {
	s1, s2 := "great", "rgeat"
	fmt.Println(isScramble(s1, s2))
	s1, s2 = "abcde", "caebd"
	fmt.Println(isScramble(s1, s2))
	s1, s2 = "a", "a"
	fmt.Println(isScramble(s1, s2))
}

func isScramble(s1 string, s2 string) bool {
	// d(i, j, k) = true if s2[j...j+k-1] is scrambled version of
	// s1[i...i+k-1]
	// for this we need to select a cut point z
	// so that s1 = a + b   and s2 = c + d
	// and then check recursively whether
	// c and d are scrambles of a and b
	// or c and d are scrambles of b and a
	// for this to be possible need the same characters in
	// a and c | b and d and so on
	// we can pre-compute this

	n := len(s1)
	if n == 1 {
		return s1 == s2
	}

	mem := map[E]bool{}
	return rec(0, 0, n, s1, s2, mem)
}

func rec(i, j, k int, s1, s2 string, mem map[E]bool) bool {
	if k == 1 {
		return s1[i] == s2[j]
	}
	if v, ok := mem[E{i, j, k}]; ok {
		return v
	}

	result := false
	for z := 1; z < k; z++ {
		rz := k - z
		a, b, c, d := i, i+z, j, j+z
		if sameChars(s1, s2, a, c, z) && sameChars(s1, s2, b, d, rz) {
			result = rec(a, c, z, s1, s2, mem) && rec(b, d, rz, s1, s2, mem)
			if result {
				break
			}
		}
		a, b, c, d = i, i+z, j+rz, j
		if sameChars(s1, s2, a, c, z) && sameChars(s1, s2, b, d, rz) {
			result = rec(a, c, z, s1, s2, mem) && rec(b, d, rz, s1, s2, mem)
			if result {
				break
			}
		}
	}

	mem[E{i, j, k}] = result
	return result
}

func sameChars(s1, s2 string, i, j, k int) bool {
	count := make([]int, 26)
	for z := 0; z < k; z++ {
		count[s1[i+z]-97] += 1
		count[s2[j+z]-97] -= 1
	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

type E struct {
	i int
	j int
	k int
}
