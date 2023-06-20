package main

import "fmt"

func main() {
	row := []int{0, 2, 1, 3}
	fmt.Println(minSwapsCouples(row))
	row = []int{3, 2, 0, 1}
	fmt.Println(minSwapsCouples(row))
}

func minSwapsCouples(row []int) int {
	// n couples, 2*n chairs in a row
	// couples are 0/1  2/3  4/5 ...
	// compute min swaps to have each couple sit next to each other
	// view each pair of positions 0/1 2/3 ... as a node in a graph
	// there is an edge between the positions of each couple, unless they are equal already
	// for a component with k nodes in this graph
	// we can use k-1 swaps to arrange them correctly
	// do this for every component
	// use dfs to find the size of components
	n := len(row) / 2
	adj := make([][]int, n)
	pos := make(map[int]int)
	for i, v := range row {
		x := i / 2
		var w int
		if v%2 == 0 {
			w = v + 1
		} else {
			w = v - 1
		}
		if y, ok := pos[w]; ok {
			if x != y {
				adj[x] = append(adj[x], y)
				adj[y] = append(adj[y], x)
			}
		}
		pos[v] = x
	}

	visited := make([]bool, n)
	result := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			result += dfs(i, -1, adj, visited) - 1
		}
	}
	return result
}

func dfs(node, parent int, adj [][]int, visited []bool) int {
	count := 1
	visited[node] = true
	for _, nb := range adj[node] {
		if nb != parent && !visited[nb] {
			count += dfs(nb, node, adj, visited)
		}
	}
	return count
}
