package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 5
	ranges := []int{3, 4, 1, 1, 0, 0}
	fmt.Println(minTaps(n, ranges))
}

func minTaps(n int, ranges []int) int {
	// greedy
	// start at 0
	// among all those that cover 0
	// choose the one that extends furthest to the right
	// then do the same thing for the next unclaimed position
	// sort the ranges by their starting point
	x := make([]Range, 0)
	for i, r := range ranges {
		if r != 0 {
			x = append(x, Range{left: i - r, right: i + r})
		}
	}
	sort.Slice(x, func(i, j int) bool {
		return x[i].left < x[j].left
	})

	curr := 0
	ri := 0
	result := 0
	for curr < n {
		maxRight := -1
		for ri < len(x) && x[ri].left <= curr {
			r := x[ri].right
			if r > maxRight {
				maxRight = r
			}
			ri += 1
		}
		if maxRight == -1 {
			return -1
		}
		result += 1
		curr = maxRight
	}

	return result
}

type Range struct {
	left  int
	right int
}
