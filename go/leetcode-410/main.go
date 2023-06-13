package main

import "fmt"

func main() {
	nums := []int{1, 4, 4}
	k := 3
	fmt.Println(splitArray(nums, k))
}

func splitArray(nums []int, k int) int {
	// split nums into k subarrays such that
	// the largest sum of any subarray is minimized
	// can use bottom-up DP
	// d(i,j) = min largest sum when splitting nums[0...i] into j parts
	// 	for j=1..k, j<=i+1
	// d(0, 1) = nums[0]

	n := len(nums)

	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = nums[i] + prefix[i-1]
	}

	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, k)
	}
	d[0][0] = nums[0]
	for i := 1; i < n; i++ {
		d[i][0] = prefix[i]
		for j := 1; j <= min(i, k-1); j++ {
			m := -1
			for z := i - 1; z >= j-1; z-- {
				l := d[z][j-1]
				r := prefix[i]
				if z >= 0 {
					r -= prefix[z]
				}
				v := max(l, r)
				if m == -1 || v < m {
					m = v
				}
			}
			d[i][j] = m
		}
	}

	fmt.Println(nums)
	fmt.Println(d)
	return d[n-1][k-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
