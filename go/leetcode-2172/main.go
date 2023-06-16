package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	numSlots := 3
	fmt.Println(maximumANDSum(nums, numSlots))
	nums = []int{1, 3, 10, 4, 7, 1}
	numSlots = 9
	fmt.Println(maximumANDSum(nums, numSlots))
	nums = []int{8, 13, 3, 15, 3, 15, 2, 15, 5, 7, 6}
	numSlots = 8
	fmt.Println(maximumANDSum(nums, numSlots))
}

func maximumANDSum(nums []int, numSlots int) int {
	// at most 18 numbers
	// at most 9 slots
	// numbers are <= 15
	// can we do dp with bitmasking?
	// can represent a subset of numbers as a length n bitmask
	// there are <= 2^18 ~ 262k of those
	// maybe d(n, k) = putting the bitmask n into k slots
	// 	where 2*k >= n
	// how do we recurse, we pick 1 or 2 for slot k+1
	// for picking one we have <18 possibilites
	// for picking 2 there at most 18*17 ~ 400
	// this might still work, can we do better?
	// maybe instead do d(n,k,c) where we have subset n, slots k and in last slot there is c == 1 or 2
	mem := map[E]int{}
	n := len(nums)
	nn := 1 << n
	// note that a slot can also be empty
	result := 0
	for c := 0; c <= 2; c++ {
		result = max(result, rec(nn-1, numSlots, c, n, nums, mem))
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// This could probably be less complicated, but such is life.
func rec(n, k, c int, count int, nums []int, mem map[E]int) int {
	if count == 0 {
		if c != 0 {
			return -1
		}
		return 0
	}
	if 2*k < count {
		return -1
	}
	if v, ok := mem[E{n, k, c}]; ok {
		return v
	}
	result := 0
	if k == 1 {
		for i := 0; i < len(nums); i++ {
			if (1<<i)&n != 0 {
				result += nums[i] & k
			}
		}
	} else if count == 1 {
		if c == 2 {
			return -1
		} else if c == 0 {
			result = max(
				rec(n, k-1, 0, count, nums, mem),
				rec(n, k-1, 1, count, nums, mem),
			)
		} else {
			for i := 0; i < len(nums); i++ {
				if (1<<i)&n != 0 {
					for z := 1; z <= k; z++ {
						result = max(result, nums[i]&z)
					}
					break
				}
			}
		}
	} else if c == 0 {
		for z := 0; z <= 2; z++ {
			x := rec(n, k-1, z, count, nums, mem)
			if x != -1 {
				result = max(result, x)
			}
		}
	} else {
		// put an element in slot k and recurse
		for i := 0; i < len(nums); i++ {
			if (1<<i)&n != 0 {
				x := rec(n^(1<<i), k, c-1, count-1, nums, mem)
				if x != -1 {
					result = max(result, x+(k&nums[i]))
				}
			}
		}
	}
	mem[E{n, k, c}] = result
	return result
}

type E struct {
	n int
	k int
	c int
}
