package main

import (
	"fmt"
	"os"
	"pp/io"
)

func main() {
	reader, err := io.NewLineReader("testcases.txt")
	defer reader.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	i := 0
	for {
		c, err := readTestCase(reader)
		if err != nil {
			if err != io.ErrorEOF {
				fmt.Println(err)
			}
			return
		}
		result := componentValue(c.nums, c.edges)
		if result == c.expected {
			fmt.Println("case", i, "ok")
		} else {
			fmt.Printf("case %v incorrect, expected %v got %v\n", i, c.expected, result)
		}
		i += 1
	}
}

type testCase struct {
	nums     []int
	edges    [][]int
	expected int
}

func readTestCase(reader *io.LineReader) (testCase, error) {
	var result testCase
	nums, err := reader.ReadIntSlice()
	if err != nil {
		return result, err
	}
	edges, err := reader.Read2DimIntSlice()
	if err != nil {
		return result, err
	}
	expected, err := reader.ReadInt()
	if err != nil {
		return result, err
	}
	result.nums = nums
	result.edges = edges
	result.expected = expected
	return result, nil
}

func componentValue(nums []int, edges [][]int) int {
	n := len(nums)
	adj := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	sum := make([]int, n)
	computeSum(0, -1, adj, nums, sum)

	// can we use the fact that the total sum isn't too large
	// and therefore can't have too many dividers?
	// total sum <= 1M -> 2^20
	result := 0
	for i := 1; i < sum[0]; i++ {
		if sum[0]%i != 0 {
			continue
		}
		r, ok := dfs(0, -1, adj, sum, i)
		if ok && r == 0 {
			result = sum[0]/i - 1
			break
		}
	}
	return result
}

func computeSum(node, parent int, adj [][]int, nums, sum []int) {
	sum[node] = nums[node]

	for _, nb := range adj[node] {
		if nb != parent {
			computeSum(nb, node, adj, nums, sum)
			sum[node] += sum[nb]
		}
	}
}

func dfs(node, parent int, adj [][]int, sum []int, target int) (int, bool) {
	if parent != -1 && len(adj[node]) == 1 {
		z := sum[node]
		if z > target {
			return 0, false
		} else if z == target {
			return 0, true
		} else {
			return z, true
		}
	}

	remaining := sum[node]
	for _, nb := range adj[node] {
		if nb != parent {
			r, ok := dfs(nb, node, adj, sum, target)
			if !ok {
				return 0, false
			}
			remaining -= sum[nb] - r
		}
	}
	if remaining == target {
		return 0, true
	} else if remaining < target {
		return remaining, true
	} else {
		return 0, false
	}
}
