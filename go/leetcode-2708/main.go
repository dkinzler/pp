package main

func maxStrength(nums []int) int64 {
	// just use all the numbers except one
	// and that is the largest negative number if there is an uneven number
	// of negative numbers
	// + alot of edge cases involving 0 and 1
	if len(nums) == 1 {
		return int64(nums[0])
	}
	maxNeg := -10
	result := 1
	countNeg := 0
	hasZero := false
	hasOne := false
	nonZero := false
	for _, v := range nums {
		if v != 0 {
			result *= v
			if v == 1 {
				hasOne = true
			}
			nonZero = true
		} else {
			hasZero = true
		}
		if v < 0 {
			countNeg += 1
			if v > maxNeg {
				maxNeg = v
			}
		}
	}
	if countNeg%2 == 1 {
		if result == maxNeg {
			if hasOne {
				return 1
			}
			if hasZero {
				return 0
			}
			// here we know we have length > 1, but if result==maxNeg
			// then there must be a 1, or at least multiple -1s
			return 1
		}
		result = result / maxNeg
		return int64(result)
	}
	if !nonZero {
		return 0
	}
	return int64(result)
}
