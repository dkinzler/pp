package main

import "fmt"

func main() {
	fmt.Println(longestAwesome("3242415"))
	fmt.Println(longestAwesome("12345678"))
	fmt.Println(longestAwesome("213123"))
	fmt.Println(longestAwesome("6037041143"))
}

func longestAwesome(s string) int {
	/*
		return length of longest substring of s, such that
		it can be made into a palindrome by swapping some characters
		a substring s can be made into a palindrome
		if at most one character appears an odd number of time
		here the characters are only digits 0-9
		what matters at the end is the parity
		i.e. whether each digit appears an even or odd number of times
		can represent this as a bit vector
		suppose we have the bit vector for range 0....i
		then we can xor it with bit vector for 0...j
		to get the parities for range j+1....i
		we can compute the min/max index for each possible bitvector
		there are only 2^10 = 1024
		then for each pair of bitvectors we try to form a solution
	*/
	n := len(s)
	k := 1 << 10
	mi := make([]int, k)
	ma := make([]int, k)
	for i := 0; i < k; i++ {
		mi[i] = -1
		ma[i] = -1
	}

	prev := 0
	for i := 0; i < n; i++ {
		curr := prev ^ (1 << (int(s[i]) - 48))
		if mi[curr] == -1 {
			mi[curr] = i
		}
		ma[curr] = i
		prev = curr
	}

	result := 1
	for i := 0; i < k; i++ {
		if isPalindrome(i) && ma[i] != -1 {
			if ma[i]+1 > result {
				result = ma[i] + 1
			}
		}
		for j := 0; j < k; j++ {
			if isPalindrome(i ^ j) {
				a, b := mi[i], ma[j]
				if a != -1 && b != -1 {
					if b-a > result {
						result = b - a
					}
				}
			}
		}
	}
	return result
}

// a bitvector represents a palindrome if it is 0
// or there is at most 1 1, i.e. it is a power of 2
func isPalindrome(b int) bool {
	if b == 0 {
		return true
	}
	if b&(b-1) == 0 {
		return true
	}
	return false
}
