package main

import (
	"fmt"
	"pp/ds/kvheap"
)

func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))
	nums1 = []int{1, 1, 2}
	nums2 = []int{1, 2, 3}
	k = 2
	fmt.Println(kSmallestPairs(nums1, nums2, k))
	nums1 = []int{1, 2}
	nums2 = []int{3}
	k = 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))
	nums1 = []int{1, 1, 2}
	nums2 = []int{1, 2, 3}
	k = 10
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	// nums1, nums2 sorted in ascending order
	// find k pairs with smallest sum
	n := len(nums1)
	result := make([][]int, 0)
	heap := kvheap.New[E, int](func(i1, i2 int) bool {
		return i1 < i2
	})
	nextIndex := make([]int, n)
	taken := 0
	right := 1
	for right <= len(nums1) {
		upper := 1 << 62
		if right < len(nums1) {
			upper = nums1[right] + nums2[0]
		}
		for j := 0; j < right; j++ {
			z := nextIndex[j]
			for z < len(nums2) {
				a, b := nums1[j], nums2[z]
				if a+b <= upper {
					heap.Update(E{j, z}, a+b)
					z += 1
				} else {
					break
				}
			}
			nextIndex[j] = z
		}
		right += 1
		if heap.Size() == 0 {
			break
		}
		for heap.Size() > 0 && taken < k {
			key, _, _ := heap.Pop()
			result = append(result, []int{nums1[key.a], nums2[key.b]})
			taken += 1
		}
		if taken == k {
			break
		}
	}

	return result
}

type E struct {
	a int
	b int
}
