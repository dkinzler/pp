package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 4
	edges := [][]int{{1, 2}, {2, 4}, {1, 3}, {2, 3}, {2, 1}}
	queries := []int{2, 3}
	fmt.Println(countPairs(n, edges, queries))
	n = 5
	edges = [][]int{{1, 5}, {1, 5}, {3, 4}, {2, 5}, {1, 3}, {5, 1}, {2, 3}, {2, 5}}
	queries = []int{1, 2, 3, 4, 5}
	fmt.Println(countPairs(n, edges, queries))
}

func countPairs(n int, edges [][]int, queries []int) []int {
	// first group edges with the same 2 nodes together
	// use just the degree of nodes to form pairs
	// then go over the edges, and for each edge
	// if we subtract from the sum of the degrees the count of that edge
	// it is smaller than the query, then remove one from the result
	// how to find the initial number of pairs?
	// use two pointers and prefix sum
	// for degree d take only those that are > d
	// as the left pointer moves further right, the right pointer moves further left
	// i.e. left pointer always increased by one
	// note that we also have to include each degree group with itself if applicable
	// to avoid counting duplicates, right pointer should be >= left
	degree := make([]int, n)
	mergedEdges := map[F]int{}
	for _, edge := range edges {
		a, b := edge[0]-1, edge[1]-1
		degree[a] += 1
		degree[b] += 1
		if a < b {
			mergedEdges[F{a, b}] += 1
		} else {
			mergedEdges[F{b, a}] += 1
		}
	}

	d := make([]E, 0)
	degToIndex := make(map[int]int)
	for _, deg := range degree {
		if i, ok := degToIndex[deg]; ok {
			d[i].count += 1
		} else {
			i := len(d)
			d = append(d, E{deg, 1})
			degToIndex[deg] = i
		}
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i].deg < d[j].deg
	})

	prefixSum := make([]int, len(d))
	prefixSum[0] = d[0].count
	for i := 1; i < len(d); i++ {
		prefixSum[i] = prefixSum[i-1] + d[i].count
	}

	result := make([]int, len(queries))
	for i, query := range queries {
		result[i] = solveQuery(mergedEdges, d, prefixSum, degree, query)
	}
	return result
}

func solveQuery(edges map[F]int, deg []E, prefixSum, nodeDegree []int, query int) int {
	m := len(deg)
	result := 0
	right := m - 1
	for left := 0; left < m; left++ {
		needed := query - deg[left].deg
		if right < left {
			right = left
		}
		for right > left && deg[right].deg > needed {
			right -= 1
		}
		ld := deg[left]
		result += ld.count * (prefixSum[m-1] - prefixSum[right])
		if ld.deg*2 > query {
			result += ld.count * (ld.count - 1) / 2
		}
	}
	for edge, count := range edges {
		a, b := edge.a, edge.b
		da, db := nodeDegree[a], nodeDegree[b]
		if da+db > query {
			if da+db-count <= query {
				result -= 1
			}
		}
	}

	return result
}

type E struct {
	deg   int
	count int
}

type F struct {
	a int
	b int
}
