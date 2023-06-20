package main

func isGoodArray(nums []int) bool {
	// if a and b have gcd(a, b) = 1
	// then there are integers x,y such that x*a + y*b = 1
	// this extends to the general case of >= 3 integers with gcd 1
	// so we can just repeatedly take the gcd
	g := nums[0]
	for _, v := range nums[1:] {
		g = gcd(g, v)
		if g == 1 {
			return true
		}
	}
	return g == 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
