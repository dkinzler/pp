package main

import "fmt"

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) []string {
	// s is short, can basically brute-force this
	// let r[i] be the set of sentences we can make with s[0...i]
	n := len(s)

	d := make(map[string]bool)
	for _, w := range wordDict {
		d[w] = true
	}

	r := make([][]string, n)
	for i := 0; i < n; i++ {
		var result []string
		if _, ok := d[s[0:i+1]]; ok {
			result = append(result, s[0:i+1])
		}
		for j := 1; j <= i; j++ {
			w := s[j : i+1]
			if _, ok := d[w]; !ok {
				continue
			}
			for _, z := range r[j-1] {
				result = append(result, z+" "+w)
			}
		}
		r[i] = result
	}
	return r[n-1]
}
