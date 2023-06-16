package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 9, 7, 3}
	fmt.Println(minimumDifference(nums))
	nums = []int{-36, 36}
	fmt.Println(minimumDifference(nums))
	nums = []int{2, -1, 0, 4, -2, -9}
	fmt.Println(minimumDifference(nums))
	nums = []int{25, 49, 39, 42, 57, 35}
	fmt.Println(minimumDifference(nums))
}

func minimumDifference(nums []int) int {
	// n is kinda small <= 15, but 2*n <= 30
	// how many subsets of size n are there at most
	// 30 over 15 -> 30!/(15!*15!) -> this is huge, too much for brute force
	// we can calculate for each possible 0/1 assignment of the first n elements (<= 2^15 ~ 32k)
	// the difference that would result, i.e. we would take for those i with 0 nums[i] and for those with 1 -1*nums[i]
	// then we group these by the number of 0s, 0...n, and sort each group
	// then we consider every possible 0/1 assignment of the last n elements
	// we again get a sum x, if there are j 0s we search among those assignments of first half with n-j 0s
	// and we want a number that is as close as possible to -x

	n2 := len(nums)
	n := n2 / 2

	d := make([][]int, n+1)
	for i := 0; i < (1 << n); i++ {
		sum := 0
		count := 0
		for j := 0; j < n; j++ {
			if (1<<j)&i == 0 {
				sum += nums[j]
				count += 1
			} else {
				sum += -1 * nums[j]
			}
		}
		d[count] = append(d[count], sum)
	}
	for i := 0; i <= n; i++ {
		sort.Ints(d[i])
	}

	result := 1 << 60
	for i := 0; i < (1 << n); i++ {
		sum := 0
		count := 0
		for j := 0; j < n; j++ {
			if (1<<j)&i == 0 {
				sum += nums[n+j]
				count += 1
			} else {
				sum += -1 * nums[n+j]
			}
		}
		v := find(d[n-count], -1*sum)
		a := abs(v + sum)
		if a < result {
			result = a
		}
	}

	return result
}

func find(x []int, target int) int {
	low := 0
	high := len(x) - 1
	result := 1 << 60
	for low <= high {
		mid := (low + high) / 2
		diff := abs(x[mid] - target)
		if x[mid] == target {
			return target
		} else if x[mid] < target {
			low = mid + 1
			if diff < abs(result-target) {
				result = x[mid]
			}
		} else {
			high = mid - 1
			if diff < abs(result-target) {
				result = x[mid]
			}
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
