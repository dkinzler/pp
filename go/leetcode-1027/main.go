package main

import "fmt"

func main() {
	nums := []int{3, 6, 9, 12}
	fmt.Println(longestArithSeqLength(nums))
	nums = []int{9, 4, 7, 2, 10}
	fmt.Println(longestArithSeqLength(nums))
	nums = []int{20, 1, 15, 3, 10, 5, 8}
	fmt.Println(longestArithSeqLength(nums))
}

func longestArithSeqLength(nums []int) int {
	// compute length of longest arithmetic subsequence of nums
	// i.e. where consecutive elements have the same difference
	// len(nums) <= 1000
	// 0 <= nums[i] <= 500
	// a sequence can be either increasing or decreasing, compute these separately
	// d(i, j) = length of longest (increasing) arithmetic sequence that ends at index i
	// and has diff j
	// to recurse need index k < i with nums[i]-nums[k] = j
	// i.e. nums[k] = nums[i] - j
	// can get this by keeping a map from values to indices

	n := len(nums)
	mi, ma := 1000, 0
	for _, v := range nums {
		if v < mi {
			mi = v
		}
		if v > ma {
			ma = v
		}
	}
	k := ma - mi

	inc := make([][]int, n)
	dec := make([][]int, n)
	for i := 0; i < n; i++ {
		inc[i] = make([]int, k+1)
		dec[i] = make([]int, k+1)
	}

	valueToIndex := map[int]int{nums[0]: 0}
	result := 1
	for j := 0; j <= k; j++ {
		inc[0][j] = 1
		dec[0][j] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= k; j++ {
			// increasing
			v := 1
			if z, ok := valueToIndex[nums[i]-j]; ok {
				v = 1 + inc[z][j]
			}
			inc[i][j] = v
			if v > result {
				result = v
			}
			// decreasing
			v = 1
			if z, ok := valueToIndex[nums[i]+j]; ok {
				v = 1 + dec[z][j]
			}
			dec[i][j] = v
			if v > result {
				result = v
			}
		}
		valueToIndex[nums[i]] = i
	}
	return result
}
