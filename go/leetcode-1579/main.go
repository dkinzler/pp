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
		result := maxNumEdgesToRemove(c.n, c.edges)
		if result == c.expected {
			fmt.Println("case", i, "ok")
		} else {
			fmt.Printf("case %v incorrect, expected %v got %v\n", i, c.expected, result)
		}
		i += 1
	}
}

type testCase struct {
	n        int
	edges    [][]int
	expected int
}

func readTestCase(reader *io.LineReader) (testCase, error) {
	var result testCase
	n, err := reader.ReadInt()
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
	result.n = n
	result.edges = edges
	result.expected = expected
	return result, nil
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	// isn't this just a mst, where we use type 3 edges first?
	// and based on that build a mst with alice/bob edges
	// but it seems to be even easier
	// we use as many type 3 edges as possible
	// this might leave us with a disconnected graph
	// then we check with dfs for both bob and alice if we can reach every node
	// if so, we can of course also build a spanning tree
	// to do this we have to use v = (n-1 - (num of type 3 edges in mst))*2
	// i.e. result is totalNumEdges - v - (num type 3 edges in mst)

	// use union find to build the mst
	d := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = i
	}

	adjA := make([][]int, n)
	adjB := make([][]int, n)

	mstSize := 0

	for _, edge := range edges {
		t, x, y := edge[0], edge[1]-1, edge[2]-1
		if t == 3 {
			a, b := find(x, d), find(y, d)
			if a != b {
				union(a, b, d)
				mstSize += 1
			}
		}

		if t == 2 || t == 3 {
			adjB[x] = append(adjB[x], y)
			adjB[y] = append(adjB[y], x)
		}
		if t == 1 || t == 3 {
			adjA[x] = append(adjA[x], y)
			adjA[y] = append(adjA[y], x)
		}
	}

	visited := make([]bool, n)
	x := dfs(0, adjA, visited)
	if x != n {
		return -1
	}
	visited = make([]bool, n)
	x = dfs(0, adjB, visited)
	if x != n {
		return -1
	}
	return len(edges) - (2*(n-1-mstSize) + mstSize)
}

func find(x int, d []int) int {
	if d[x] == x {
		return x
	}
	d[x] = find(d[x], d)
	return d[x]
}

func union(a, b int, d []int) {
	d[a] = b
}

func dfs(node int, adj [][]int, visited []bool) int {
	if visited[node] {
		return 0
	}

	visited[node] = true
	v := 1
	for _, nb := range adj[node] {
		v += dfs(nb, adj, visited)
	}
	return v
}
