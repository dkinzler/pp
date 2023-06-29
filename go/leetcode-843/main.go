package main

// This is the Master's API interface.
// You should not implement it, or speculate about its implementation
type Master struct{}

func (this *Master) Guess(word string) int {
	return 0
}

func findSecretWord(words []string, master *Master) {
	n := len(words)
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			d[i][j] = same(words[i], words[j])
		}
	}

	out := make([]bool, n)
	for i := 0; i < n; i++ {
		if !out[i] {
			curr := i
			for {
				correct := master.Guess(words[curr])
				if correct == 6 {
					return
				}
				out[curr] = true
				next := -1
				for j := 0; j < n; j++ {
					if !out[j] {
						if d[curr][j] == correct {
							next = j
						} else {
							out[j] = true
						}
					}
				}
				if next == -1 {
					break
				} else {
					curr = next
				}
			}
		}
	}
}

func same(s, t string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			result += 1
		}
	}
	return result
}
