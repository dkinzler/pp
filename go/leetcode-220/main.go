package main

import (
	"fmt"
	"pp/ds/kvheap"
	"sort"
)

func main() {
	nums, indexDiff, valueDiff := []int{1, 2, 3, 1}, 3, 0
	fmt.Println(containsNearbyAlmostDuplicate(nums, indexDiff, valueDiff))
	nums, indexDiff, valueDiff = []int{1, 5, 9, 1, 5, 9}, 2, 3
	fmt.Println(containsNearbyAlmostDuplicate(nums, indexDiff, valueDiff))
	nums, indexDiff, valueDiff = []int{10, 100, 11, 9}, 1, 2
	fmt.Println(containsNearbyAlmostDuplicate(nums, indexDiff, valueDiff))
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	// one solution is to use a sliding window of the last indexDiff elements
	// keep them sorted using an ordered set
	// can then easily check if there is a value within valueDiff of current value

	// is there an easier solution?
	// can also turn the problem around
	// sort by value, keep those within valueDiff in the sliding window
	// and then check if there is one with a suitable index
	// can we group indices into indexDiff windows
	// <= n such windows
	// i.e. 0....indexDiff-1, indexDiff, ...2*indexDiff-1, ....
	// suppose our current element as index j
	// then if the bucket that j belongs in already has an element we return true
	// if not we have to check the two adjacent buckets
	// for the bucket on the right we want the smallest element
	// and those on the left we want the biggest one
	// so for each bucket need to keep track of min and max
	n := len(nums)
	x := make([]E, n)
	for i, v := range nums {
		x[i] = E{v: v, i: i}
	}

	sort.Slice(x, func(i, j int) bool {
		return x[i].v < x[j].v
	})

	nBuckets := (n-1)/indexDiff + 1
	// could do min and max heap for each bucket
	buckets := make([]*Bucket, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = NewBucket()
	}

	left := 0
	for i, e := range x {
		bi := e.i / indexDiff
		b := buckets[bi]
		if i == 0 {
			b.min.Update(e.i, e.i)
			b.max.Update(e.i, e.i)
			b.count = 1
		} else {
			for left < i {
				if e.v-x[left].v <= valueDiff {
					break
				}
				li := x[left].i
				bli := li / indexDiff
				bl := buckets[bli]
				bl.min.RemoveKey(li)
				bl.max.RemoveKey(li)
				bl.count -= 1
				left += 1
			}

			if b.count > 0 {
				return true
			}
			if bi > 0 {
				leftBucket := buckets[bi-1]
				leftMax, _, ok := leftBucket.max.First()
				if ok && e.i-leftMax <= indexDiff {
					return true
				}
			}
			if bi < nBuckets-1 {
				rightBucket := buckets[bi+1]
				rightMin, _, ok := rightBucket.min.First()
				if ok && rightMin-e.i <= indexDiff {
					return true
				}
			}
			b.min.Update(e.i, e.i)
			b.max.Update(e.i, e.i)
			b.count += 1
		}
	}

	return false
}

type E struct {
	v int
	i int
}

type Bucket struct {
	min   *kvheap.KeyValueHeap[int, int]
	max   *kvheap.KeyValueHeap[int, int]
	count int
}

func NewBucket() *Bucket {
	return &Bucket{
		count: 0,
		min:   kvheap.New[int, int](func(i1, i2 int) bool { return i1 < i2 }),
		max:   kvheap.New[int, int](func(i1, i2 int) bool { return i1 > i2 }),
	}
}
