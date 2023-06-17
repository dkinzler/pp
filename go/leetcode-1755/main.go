package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, -7, 3, 5}
	goal := 6
	fmt.Println(minAbsDifference(nums, goal))
	nums = []int{7, -9, 15, -2}
	goal = -5
	fmt.Println(minAbsDifference(nums, goal))
	nums = []int{1, 2, 3}
	goal = -7
	fmt.Println(minAbsDifference(nums, goal))
	nums = []int{1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}
	goal = 45
	fmt.Println(minAbsDifference(nums, goal))
	nums = []int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2}
	goal = -3
	fmt.Println(minAbsDifference(nums, goal))
}

func minAbsDifference(nums []int, goal int) int {
	// for first half compute every possible sum <= 2^20 ~1M
	// sort them
	// for second half compute every possible sum and find optimal in first half with binary search to combine
	// to create as close as possible to goal
	n := len(nums)
	if n <= 20 {
		result := abs(goal)
		nn := 1 << n
		x := make([]int, nn)
		for i := 1; i < nn; i++ {
			sum := 0
			for j := 0; j < n; j++ {
				if (1<<j)&i != 0 {
					sum = nums[j] + x[i^(1<<j)]
					break
				}
			}
			x[i] = sum
			v := abs(sum - goal)
			if v < result {
				result = v
			}
		}
		return result
	}

	k := n / 2
	result := abs(goal)
	x := make([]int, 1<<k)
	for i := 1; i < (1 << k); i++ {
		sum := 0
		for j := 0; j < k; j++ {
			if (1<<j)&i != 0 {
				sum = nums[j] + x[i^(1<<j)]
				break
			}
		}
		x[i] = sum
	}

	sort.Ints(x)

	y := make([]int, 1<<(n-k))
	for i := 1; i < (1 << (n - k)); i++ {
		sum := 0
		for j := 0; j < n-k; j++ {
			if (1<<j)&i != 0 {
				sum = nums[k+j] + y[i^(1<<j)]
				break
			}
		}
		y[i] = sum
		x := search(x, -1*(sum-goal))
		if x < result {
			result = x
		}
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

// find number such that abs(number+target) is minimal
func search(x []int, target int) int {
	result := 1 << 60
	low := 0
	high := len(x) - 1
	for low <= high {
		mid := (low + high) / 2
		v := x[mid]
		if v == target {
			return 0
		} else if v < target {
			low = mid + 1
			if abs(v-target) < result {
				result = abs(v - target)
			}
		} else {
			high = mid - 1
			if abs(v-target) < result {
				result = abs(v - target)
			}
		}
	}
	return result
}
