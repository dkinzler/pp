package main

func findValidSplit(nums []int) int {
	// could use a sliding window approach where we keep track of
	// the unique prime factors and their frequency on the left and right
	// if there is no overlap, the products are co-prime
	max := 0
	for _, x := range nums {
		if x > max {
			max = x
		}
	}

	factors := make([]int, max+1)
	for i := 2; i <= max; i++ {
		if factors[i] != 0 {
			continue
		}
		factors[i] = i
		for j := i; i*j <= max; j++ {
			if factors[j*i] == 0 {
				factors[j*i] = i
			}
		}
	}

	n := len(nums)

	rightPrimes := make(map[int]int)
	ff := make([][]int, n)

	for i, x := range nums {
		ff[i] = factor(x, factors)
		for _, p := range ff[i] {
			rightPrimes[p] += 1
		}
	}

	leftPrimes := make(map[int]int)
	intersectCount := 0
	for i := 0; i < n-1; i++ {
		// move number from right to left side
		for _, p := range ff[i] {
			v := rightPrimes[p]
			rightPrimes[p] -= 1
			leftPrimes[p] += 1
			if leftPrimes[p] == 1 {
				if v > 1 {
					intersectCount += 1
				}
			} else if v == 1 {
				intersectCount -= 1
			}
		}
		if intersectCount == 0 {
			return i
		}
	}
	return -1
}

func factor(n int, factors []int) []int {
	var result []int
	last := -1
	for n > 1 {
		curr := factors[n]
		if curr != last {
			last = curr
			result = append(result, curr)
		}
		n = n / factors[n]
	}
	return result
}
