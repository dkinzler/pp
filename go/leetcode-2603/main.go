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
		result := collectTheCoins(c.coins, c.edges)
		if result == c.expected {
			fmt.Println("case", i, "ok")
		} else {
			fmt.Printf("case %v incorrect, expected %v got %v\n", i, c.expected, result)
		}
		i += 1
	}
}

type testCase struct {
	coins    []int
	edges    [][]int
	expected int
}

func readTestCase(reader *io.LineReader) (testCase, error) {
	var result testCase
	coins, err := reader.ReadIntSlice()
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
	result.coins = coins
	result.edges = edges
	result.expected = expected
	return result, nil
}

func collectTheCoins(coins []int, edges [][]int) int {
	n := len(coins)
	adj := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	d := make([]int, n)
	e := make([]int, n)
	dfs(0, -1, coins, adj, e, d)
	pp := -1
	if coins[0] == 1 {
		pp = 0
	}
	return dfs2(0, -1, 0, pp, adj, e, d, coins)
}

// suppose we do dfs from a node
func dfs(node, parent int, coins []int, adj [][]int, e, d []int) {
	if parent != -1 && len(adj[node]) == 1 {
		e[node] = 0
		if coins[node] == 1 {
			d[node] = 0
		} else {
			d[node] = -1
		}
		return
	}

	re := 0
	rd := -1
	if coins[node] == 1 {
		rd = 0
	}
	for _, nb := range adj[node] {
		if nb == parent {
			continue
		}
		dfs(nb, node, coins, adj, e, d)
		if d[nb] != -1 {
			z := d[nb] + 1
			if rd == -1 || z > rd {
				rd = z
			}
			if z >= 3 {
				re += 2
			}
		}
		re += e[nb]
	}
	e[node] = re
	d[node] = rd
}

func dfs2(node, parent, parentEdges, parentMaxDist int, adj [][]int, e, d, coins []int) int {
	if parent != -1 && len(adj[node]) == 1 {
		result := 0
		if parentEdges > 0 {
			result = parentEdges + 2
		}
		if parentMaxDist >= 2 {
			result = parentEdges + 2
		}
		return result
	}

	maxDist1 := -1
	maxDist2 := -1
	for _, nb := range adj[node] {
		if nb == parent {
			continue
		}
		z := d[nb]
		if z != -1 {
			if maxDist1 == -1 {
				maxDist1 = z
			} else if z > maxDist1 {
				maxDist2 = maxDist1
				maxDist1 = z
			} else if maxDist2 == -1 || z > maxDist2 {
				maxDist2 = z
			}
		}
	}

	result := e[node] + parentEdges
	if parentEdges > 0 {
		result += 2
	} else if parentMaxDist >= 3 {
		result += 2
	}

	baseEdges := parentEdges
	if parentMaxDist >= 3 {
		baseEdges += 2
	}
	baseDist := -1
	if parentMaxDist != -1 {
		baseDist = parentMaxDist + 1
	}
	if coins[node] == 1 {
		if baseDist == -1 {
			baseDist = 1
		}
	}
	for _, nb := range adj[node] {
		if nb == parent {
			continue
		}
		edges := baseEdges + e[node] - e[nb]
		if d[nb] >= 2 {
			edges -= 2
		}
		maxDist := baseDist
		if d[nb] == maxDist1 {
			if maxDist2 != -1 {
				if maxDist == -1 || maxDist2+2 > maxDist {
					maxDist = maxDist2 + 2
				}
			}
		} else {
			if maxDist1 != -1 {
				if maxDist == -1 || maxDist1+2 > maxDist {
					maxDist = maxDist1 + 2
				}
			}
		}

		x := dfs2(nb, node, edges, maxDist, adj, e, d, coins)
		if result == -1 || x < result {
			result = x
		}
	}

	return result
}
