package main

func minimumDeletions(s string) int {
	// s consists only of characters a/b
	// compute min number of deletions needed to make s balanced
	// s is balanced if no b appears before an a
	// i.e. first all a then b
	// can iterate over the possible choices of cut point
	// iterate over string, at every a we can assume it is the last one
	// i.e. we have to delete all b's to the left and all a's to the right
	n := len(s)
	aCount := 0
	bCount := 0
	for i := 0; i < n; i++ {
		if s[i] == byte('a') {
			aCount += 1
		} else {
			bCount += 1
		}
	}

	result := aCount
	if bCount < result {
		result = bCount
	}
	currA := 0
	currB := 0
	for i := 0; i < n; i++ {
		if s[i] == byte('a') {
			currA += 1
			x := currB + aCount - currA
			if x < result {
				result = x
			}
		} else {
			currB += 1
		}
	}

	return result
}
