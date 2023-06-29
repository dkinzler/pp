package main

func lengthOfLIS(nums []int) int {
	// longest strictly increasing subsequence
	// O(n^2) is easy
	// for each i compute d(i) = longest subsequence ending at index i
	// look at each j < i to use as previous element

	// return simpleLIS(nums)

	// but we can do better
	// suppose the longest sequence so far ends in value v and has length l
	// and at index i we have value w and length k
	// if v < w then we can use that sequence to get a higher value
	// if v >= w there are two cases
	//    l =< k, i.e. then we don't need (v, l), since any later subsequence
	//    	can just use (w, k)
	//    l > k keep (v,l)
	// the point is that this will form an increasing sequence in both
	// the value and length of sequence, and thus we can do binary search
	// just need to keep these as a sorted list
	// in summary: when we add (w, k) we can remove all (v, l)
	// with v >= w and l <= k
	// but every index will be added/removed at most once

	return betterLIS(nums)
}

func simpleLIS(nums []int) int {
	n := len(nums)
	d := make([]int, n)
	d[0] = 1
	result := 1
	for i := 1; i < n; i++ {
		v := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && d[j]+1 > v {
				v = d[j] + 1
			}
		}
		d[i] = v
		if v > result {
			result = v
		}
	}
	return result
}

func betterLIS(nums []int) int {
	return 0
}
