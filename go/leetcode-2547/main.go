package main

func minCost(nums []int, k int) int {
	// split nums into subarrays such that
	// sum of importance of subarrays is minimized
	// importance = k + length of subarray with all elements of frequency 1 removed
	// can recursively compute minCost for range j....n-1
	// can do this top down with memoization
	// how do we efficiently compute the trimmed length of a subarray
	// simply iterating will make the algorithm too slow for n =~ 1000
	// note that numbers are < n
	// could precompute these values
	// trimmed length for all i,j
	// and there we could keep track of numbers, yes yes

	n := len(nums)

	tl := make([][]int, n)
	for i := 0; i < n; i++ {
		tl[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		freq := make(map[int]int)
		length := 0
		for j := i; j < n; j++ {
			freq[nums[j]] += 1
			v := freq[nums[j]]
			if v == 2 {
				length += 2
			} else if v > 2 {
				length += 1
			}
			tl[i][j] = length
		}
	}

	mem := map[int]int{}
	return rec(0, nums, k, mem, tl)
}

func rec(start int, nums []int, k int, mem map[int]int, tl [][]int) int {
	if start >= len(nums) {
		return 0
	}
	if start == len(nums)-1 {
		return k
	}

	if v, ok := mem[start]; ok {
		return v
	}

	result := -1
	for j := start; j < len(nums); j++ {
		x := k + tl[start][j] + rec(j+1, nums, k, mem, tl)
		if result == -1 || x < result {
			result = x
		}
	}
	mem[start] = result
	return result
}
