package main

func countTriplets(nums []int) int {
	// count triplets i, j, k such that nums[i]&nums[j]&nums[k] == 0
	// i,j,k do not have to be distinct
	// the numbers are < 2^16

	nn := 1 << 16

	d := make([]int, nn)
	for _, v := range nums {
		for _, w := range nums {
			d[v&w] += 1
		}
	}

	result := 0
	for v, count := range d {
		if count > 0 {
			for _, w := range nums {
				if v&w == 0 {
					result += count
				}
			}
		}
	}
	return result
}
