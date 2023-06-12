package main

import "sort"

func earliestFullBloom(plantTime []int, growTime []int) int {
	// note that we want to plant every day, so the total time planting is always the same
	// i.e. we will always finish planting on the same day
	// and we can plant on seed after another, mixing should not help
	// we probably want those that grow for a long time at the start
	// they day we are finished planting is always the same, so having a long grow last does not make sense
	// what if there are two seeds that have the same grow time
	// does the order matter? no ->
	n := len(plantTime)

	x := make([]E, n)
	for i := 0; i < n; i++ {
		x[i] = E{pt: plantTime[i], gt: growTime[i]}
	}

	sort.Slice(x, func(i, j int) bool {
		return x[i].gt > x[j].gt
	})

	startPlant := 0
	result := 0
	for _, e := range x {
		startPlant += e.pt
		result = max(result, startPlant+e.gt)
	}

	return result
}

type E struct {
	pt int
	gt int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
