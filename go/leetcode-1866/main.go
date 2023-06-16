package main

import "fmt"

func main() {
	n, k := 3, 2
	fmt.Println(rearrangeSticks(n, k))
	n, k = 5, 5
	fmt.Println(rearrangeSticks(n, k))
	n, k = 20, 11
	fmt.Println(rearrangeSticks(n, k))
}

func rearrangeSticks(n int, k int) int {
	// dp and combinatorics
	// d(i, j) = number of arrangements for i sticks where j are visible
	//     1 <= j <= i
	// think about where to put the smallest element
	// d(i,j) = d(i-1, j)*(i-1) because we can place the 1
	// after every stick and it will not be visible
	//  + d(i-1, j-1) where we place it at the front and it is visible
	// lets use top down dp with memoization
	mem := map[E]int{}
	result := rec(n, k, mem)
	// fmt.Println(mem)
	return result
}

func rec(n, k int, mem map[E]int) int {
	if n < k || k < 1 {
		return 0
	}
	if n == 1 || n == k {
		return 1
	}

	if v, ok := mem[E{n, k}]; ok {
		return v
	}

	mod := 1000000007
	if k == 1 {
		// largest one at the front, arrangement of rest doesn't matter
		// so (n-1)! ways
		result := 1
		for i := 2; i <= n-1; i++ {
			result = (result * i) % mod
		}
		mem[E{n, k}] = result
		return result
	}

	result := (rec(n-1, k-1, mem) + rec(n-1, k, mem)*(n-1)) % mod
	mem[E{n, k}] = result
	return result
}

type E struct {
	n int
	k int
}
