package main

import "sort"

func minCost(nums []int, cost []int) int64 {
	n := len(nums)
	x := make([]E, n)
	for i, num := range nums {
		x[i] = E{int64(num), int64(cost[i])}
	}

	sort.Slice(x, func(i, j int) bool {
		return x[i].num < x[j].num
	})

	costPrefix := make([]int64, n)
	costPrefix[0] = x[0].cost
	for i := 1; i < n; i++ {
		costPrefix[i] = costPrefix[i-1] + x[i].cost
	}

	var curr int64
	for _, e := range x {
		curr += (e.num - x[0].num) * e.cost
	}
	result := curr

	for i := 1; i < n; i++ {
		e := x[i]
		// curr target is now e.num
		diff := e.num - x[i-1].num
		curr += diff * costPrefix[i-1]
		curr -= diff * (costPrefix[n-1] - costPrefix[i-1])
		if curr < result {
			result = curr
		}
	}

	return result
}

type E struct {
	num  int64
	cost int64
}
