package main

import (
	"sort"
)

func minWastedSpace(packages []int, boxes [][]int) int {
	// compute the min wasted space for each supplier
	// do this by sorting the packages and box sizes and then iterating through both one by one
	// for each package we want to choose the smallest box that is equal or larger
	sort.Ints(packages)
	mod := 1000000007

	n := len(packages)
	packagePrefix := make([]int, n)
	packagePrefix[0] = packages[0]
	for i := 1; i < n; i++ {
		packagePrefix[i] = packages[i] + packagePrefix[i-1]
	}

	result := -1
	for _, b := range boxes {
		sort.Ints(b)
		x := computeWastedSpace(packages, packagePrefix, b)
		if x != -1 {
			if result == -1 || x < result {
				result = x
			}
		}
	}
	return result % mod
}

func computeWastedSpace(packages []int, packagePrefix []int, boxes []int) int {
	ws := 0
	pi := 0
	for _, box := range boxes {
		i := findIndex(packages, box)
		if i == -1 || i < pi {
			continue
		}
		s := packagePrefix[i]
		if pi > 0 {
			s -= packagePrefix[pi-1]
		}

		// 5 4 3
		count := i - pi + 1
		ws = ws + count*box - s

		pi = i + 1
		if pi >= len(packages) {
			break
		}
	}
	if pi < len(packages) {
		return -1
	}
	return ws
}

// return largest index i such that packages[i] <= box
// or -1 if no such index exists
func findIndex(packages []int, box int) int {
	result := -1
	low := 0
	high := len(packages) - 1
	for low <= high {
		mid := (low + high) / 2
		if packages[mid] <= box {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}
