package main

import (
	"fmt"
	"pp/ds/kvheap"
)

func main() {
	costs := []int{17, 12, 10, 2, 7, 2, 11, 20, 8}
	k := 3
	candidates := 4
	fmt.Println(totalCost(costs, k, candidates))
	costs = []int{1, 2, 4, 1}
	k = 3
	candidates = 3
	fmt.Println(totalCost(costs, k, candidates))
	costs = []int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}
	k = 11
	candidates = 2
	fmt.Println(totalCost(costs, k, candidates))
}

func totalCost(costs []int, k int, candidates int) int64 {
	// costs[i] = cost to hire i-th worker
	// want to choose k workers
	// do k iterations where in each iteration
	// choose the worker with lowest cost (break ties by index) among the first and last candidates
	// workers
	// keep a heap that contains the current elements
	// at each step we pop the min and add the next to inner elements
	n := len(costs)
	heap := kvheap.New[int, Value](func(v1, v2 Value) bool {
		if v1.cost < v2.cost {
			return true
		} else if v1.cost > v2.cost {
			return false
		} else {
			return v1.index < v2.index
		}
	})
	for i := 0; i < candidates; i++ {
		if i >= n {
			break
		}
		heap.Update(i, Value{costs[i], i})
	}
	for i := n - candidates; i < n; i++ {
		if i <= candidates-1 {
			continue
		}
		heap.Update(i, Value{costs[i], i})
	}
	left := candidates - 1
	right := n - candidates
	var result int64
	for i := 0; i < k; i++ {
		_, v, _ := heap.Pop()
		result += int64(v.cost)
		if left < right {
			if v.index <= left {
				heap.Update(left+1, Value{costs[left+1], left + 1})
				left += 1
			} else {
				heap.Update(right-1, Value{costs[right-1], right - 1})
				right -= 1
			}
		}
	}

	return result
}

type Value struct {
	cost  int
	index int
}
