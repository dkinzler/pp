package main

import "fmt"

func main() {
	nums := []int{1, 1, 1}
	fmt.Println(waysToSplit(nums))
	nums = []int{1, 2, 2, 2, 5, 0}
	fmt.Println(waysToSplit(nums))
	nums = []int{3, 2, 1}
	fmt.Println(waysToSplit(nums))
}

func waysToSplit(nums []int) int {
	n := len(nums)

	prefixSum := make([]int, n)
	prefixSum[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	mod := 1000000007
	result := 0
	for i := 0; i < n-2; i++ {
		// calculate number of ways where first array ends at i, i.e. i is still included
		// need to find first and last positions where second cut is possible
		// need to find first and last positions where second cut is possible
		left := search(i, prefixSum, true)
		right := search(i, prefixSum, false)
		if left != -1 {
			result = (result + right - left + 1) % mod
		}
	}
	return result
}

func search(i int, prefixSum []int, findLowest bool) int {
	firstSum := prefixSum[i]
	low := i + 1
	high := len(prefixSum) - 2
	result := -1
	for low <= high {
		mid := (low + high) / 2
		secondSum := sum(i+1, mid, prefixSum)
		thirdSum := sum(mid+1, len(prefixSum)-1, prefixSum)
		if firstSum > secondSum {
			// too far left
			low = mid + 1
		} else if secondSum > thirdSum {
			// too far right
			high = mid - 1
		} else {
			result = mid
			if findLowest {
				// find lowest
				high = mid - 1
			} else {
				// find highest
				low = mid + 1
			}
		}
	}
	return result
}

func sum(i, j int, prefixSum []int) int {
	result := prefixSum[j]
	if i > 0 {
		result -= prefixSum[i-1]
	}
	return result
}
