package main

import "fmt"

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println(findRepeatedDnaSequences(s))
	s = "AAAAAAAAAAAAA"
	fmt.Println(findRepeatedDnaSequences(s))
}

func findRepeatedDnaSequences(s string) []string {
	// we could do string hashing, but since the substrings we want to
	// consider all have length 10, and there are only 4 different characters
	// we can use bit strings
	z := 0
	seen := map[int]int{}
	for i := 0; i < len(s); i++ {
		z = z << 2
		z = z &^ (3 << 20)
		switch s[i] {
		case byte('A'):
			z |= 1
		case byte('C'):
			z |= 2
		case byte('G'):
			z |= 3
		default:
		}
		if i >= 9 {
			seen[z] += 1
		}
	}

	result := make([]string, 0)
	for k, v := range seen {
		if v > 1 {
			result = append(result, bitmaskToString(k))
		}
	}
	return result
}

func bitmaskToString(mask int) string {
	result := make([]byte, 10)
	for i := 0; i < 10; i++ {
		switch mask & 3 {
		case 0:
			result[9-i] = 'T'
		case 1:
			result[9-i] = 'A'
		case 2:
			result[9-i] = 'C'
		case 3:
			result[9-i] = 'G'
		}
		mask = mask >> 2
	}
	return string(result)
}
