package main

import "fmt"

func main() {
	cards := []int{4, 1, 8, 7}
	fmt.Println(judgePoint24(cards))
	cards = []int{1, 2, 1, 2}
	fmt.Println(judgePoint24(cards))
	cards = []int{3, 3, 8, 8}
	fmt.Println(judgePoint24(cards))
}

func judgePoint24(cards []int) bool {
	// 4 cards with values 1-9
	// can use numbers in any order
	// try to create a mathematical expression
	// using +, -, *, / that results in 24
	// return true if this is possible, false otherwise
	// we can probably just brute force this
	// try every order of numbers
	// and for each i<j compute the possible values
	epsilon := 0.00001
	for _, c := range permutations(cards) {
		mem := make(map[E][]float64)
		for _, v := range rec(0, 3, c, mem) {
			if v >= 24-epsilon && v <= 24+epsilon {
				return true
			}
		}
	}

	return false
}

func permutations(cards []int) [][]float64 {
	result := make([][]float64, 0)
	for i := 0; i < (1 << 8); i++ {
		indices := make([]int, 4)
		for j := 0; j < 4; j++ {
			index := (i >> (j * 2)) & 3
			indices[j] = index
		}
		allDiff := true
		for j := 0; j < 4; j++ {
			for k := j + 1; k < 4; k++ {
				if indices[j] == indices[k] {
					allDiff = false
					break
				}
			}
			if !allDiff {
				break
			}
		}
		if allDiff {
			x := make([]float64, 4)
			for j, v := range indices {
				x[j] = float64(cards[v])
			}
			result = append(result, x)
		}
	}
	return result
}

func rec(start, end int, cards []float64, mem map[E][]float64) []float64 {
	if start == end {
		return []float64{cards[start]}
	}

	if v, ok := mem[E{start, end}]; ok {
		return v
	}

	seen := map[float64]bool{}
	var result []float64
	// split calc into 2 parts
	for i := start; i < end; i++ {
		x1 := rec(start, i, cards, mem)
		x2 := rec(i+1, end, cards, mem)
		for _, v1 := range x1 {
			for _, v2 := range x2 {
				z := v1 + v2
				if _, ok := seen[z]; !ok {
					result = append(result, z)
					seen[z] = true
				}
				z = v1 - v2
				if _, ok := seen[z]; !ok {
					result = append(result, z)
					seen[z] = true
				}
				z = v1 * v2
				if _, ok := seen[z]; !ok {
					result = append(result, z)
					seen[z] = true
				}
				if v2 != 0 {
					z = v1 / v2
					if _, ok := seen[z]; !ok {
						result = append(result, z)
						seen[z] = true
					}
				}
			}
		}
	}

	mem[E{start, end}] = result
	return result
}

type E struct {
	start int
	end   int
}
