package main

import "fmt"

func main() {
	nums := []int{1, 3, 5}
	fmt.Println(findMin(nums))
	nums = []int{2, 2, 2, 0, 1}
	fmt.Println(findMin(nums))
	nums = []int{2, 2, 2, 2, 1}
	fmt.Println(findMin(nums))
	nums = []int{4, 5, 6, 7, 1, 1, 2, 3}
	fmt.Println(findMin(nums))
	nums = []int{10, 1, 10, 10, 10}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
	// of course O(n) is not good enough here
	// we want to find a position i such that a[i-1] > a[i] <= a[i+1]
	// if all elements are equal, this might not exist
	// can we do binary search?
	// just pick an element in the middle and compare with start of interval
	// i.e. in beginning pick x = a[n-1/2] and compare with a[0] and a[n-1]
	// suppose a[0] > x, use x as current min, min must be in range 1,.....(n-1)/2-1
	// if a[0] <= x, set min to a[0], min must be in range (n-1)/2+1,. ... n-1
	// if all elements are equal, we will always need O(n)
	n := len(nums) - 1

	result := 1 << 60
	left := 0
	right := n
	for left <= right {
		mid := (left + right) / 2
		x := nums[mid]
		if nums[left] > x {
			result = min(result, x)
			left += 1
			right = mid - 1
		} else if nums[left] < x {
			result = min(result, nums[left])
			left = mid + 1
		} else {
			if nums[right] < x {
				result = min(result, nums[right])
				left = mid + 1
				right -= 1
			} else {
				result = min(result, x)
				left += 1
				right -= 1
			}
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
