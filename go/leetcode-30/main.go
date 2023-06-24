package main

import "fmt"

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	fmt.Println(findSubstring(s, words))
}

func findSubstring(s string, words []string) []int {
	// first there can be duplicates in words
	// so transform it into a set of (word, count)

	// for each pos in s we can calculate if the string starting there is in words, since all the words have same length
	// can then iterate over each starting point
	// but we can do a sliding window for each of the positions
	n := len(s)
	m := len(words)
	k := len(words[0])

	nextIndex := 0
	wordToIndex := make(map[string]int)
	for _, word := range words {
		if _, ok := wordToIndex[word]; !ok {
			wordToIndex[word] = nextIndex
			nextIndex += 1
		}
	}
	ndiff := len(wordToIndex)
	count := make([]int, ndiff)
	for _, word := range words {
		count[wordToIndex[word]] += 1
	}

	d := make([]int, n)
	for i := 0; i < n; i++ {
		if i+k-1 >= n {
			d[i] = -1
		} else if v, ok := wordToIndex[s[i:i+k]]; ok {
			d[i] = v
		} else {
			d[i] = -1
		}
	}

	result := make([]int, 0)
	for i := 0; i < k; i++ {
		left := i
		right := i
		cc := make([]int, ndiff)
		correct := 0
		size := 1
		for right+k-1 < n {
			wi := d[right]
			if wi != -1 {
				cc[wi] += 1
				if cc[wi] == count[wi] {
					correct += 1
				} else if cc[wi] == count[wi]+1 {
					correct -= 1
				}
			}
			for size > m {
				wi = d[left]
				if wi != -1 {
					if cc[wi] == count[wi] {
						correct -= 1
					} else if cc[wi] == count[wi]+1 {
						correct += 1
					}
					cc[wi] -= 1
				}
				left += k
				size -= 1
			}
			if correct == ndiff {
				result = append(result, left)
			}
			right += k
			size += 1
		}
	}
	return result
}
