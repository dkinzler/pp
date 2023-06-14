package main

import (
	"pp/ds/kvheap"
	"sort"
)

func minInterval(intervals [][]int, queries []int) []int {
	// could we keep a heap?
	// sort the queries and intervals
	// move through the queries 1 by 1
	// add all the intervals that contain the current number
	// and remove those that don't anymore, but how do we find those?
	//		could use a second heap that is sorted by endpoint, yes yes
	n := len(intervals)
	x := make([]E, n)
	for i, inv := range intervals {
		x[i] = E{i, inv[0], inv[1]}
	}
	sort.Slice(x, func(i, j int) bool {
		return x[i].l < x[j].l
	})

	sortedQueries := make([]Q, len(queries))
	for i, q := range queries {
		sortedQueries[i] = Q{i, q}
	}
	sort.Slice(sortedQueries, func(i, j int) bool {
		return sortedQueries[i].q < sortedQueries[j].q
	})

	heapBySize := kvheap.New[int, int](func(i1, i2 int) bool { return i1 < i2 })
	heapByEndpoint := kvheap.New[int, int](func(i1, i2 int) bool { return i1 < i2 })
	left := 0
	result := make([]int, len(queries))
	for _, q := range sortedQueries {
		for left < n && x[left].l <= q.q {
			inv := x[left]
			heapBySize.Update(inv.i, inv.r-inv.l+1)
			heapByEndpoint.Update(inv.i, inv.r)
			left += 1
		}
		for heapByEndpoint.Size() > 0 {
			k, v, _ := heapByEndpoint.First()
			if v < q.q {
				heapByEndpoint.Pop()
				heapBySize.RemoveKey(k)
			} else {
				break
			}
		}

		if heapBySize.Size() > 0 {
			_, v, _ := heapBySize.First()
			result[q.i] = v
		} else {
			result[q.i] = -1
		}
	}

	return result
}

type E struct {
	i int
	l int
	r int
}

type Q struct {
	i int
	q int
}
