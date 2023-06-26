package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 2, 3, 1}
	fmt.Println(reversePairs(nums))
	nums = []int{2, 4, 3, 5, 1}
	fmt.Println(reversePairs(nums))
	nums = []int{1, 3, 2, 5, 6, 8}
	fmt.Println(reversePairs(nums))
	nums = []int{-5, -5}
	fmt.Println(reversePairs(nums))
}

func reversePairs(nums []int) int {
	// find number of pairs i < j such that nums[i] > 2*nums[j]
	// sort the numbers, then do a sliding window with a BIT
	// when we are at index i in the sorted array
	// add all original indices (in the unsorted array)
	// of those j < i such that nums[i] > 2*nums[j] to the BIT
	n := len(nums)
	x := make([]E, n)
	for i, v := range nums {
		x[i] = E{i, v}
	}
	sort.Slice(x, func(i, j int) bool {
		a, b := x[i].v, x[j].v
		if a < b {
			return true
		} else if a > b {
			return false
		} else {
			return x[i].i > x[j].i
		}
	})

	bit := NewBIT(n)
	left := 0
	result := 0
	for i := 0; i < n; i++ {
		for left < n && 2*x[left].v < x[i].v {
			bit.Update(x[left].i, 1)
			left += 1
		}
		result += bit.GetSum(n-1) - bit.GetSum(x[i].i)
	}
	return result
}

type E struct {
	i int
	v int
}

type BIT []int

func NewBIT(n int) BIT {
	return make(BIT, n+1)
}

func (b BIT) Update(i, v int) {
	i += 1
	for i < len(b) {
		b[i] += v
		i += i & (-i)
	}
}

func (b BIT) GetSum(i int) int {
	result := 0
	i = i + 1
	for i > 0 {
		result += b[i]
		i -= i & (-i)
	}
	return result
}
