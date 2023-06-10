package main

func minTrioDegree(n int, edges [][]int) int {
	// can't just brute force this?
	// n<=400
	// if for each node we store the degree and keep a map of neighbors for each node
	// then we can check if we have a trio and compute the degree in O(1)

	adj := make([][]bool, n)
	deg := make([]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]bool, n)
	}
	for _, edge := range edges {
		a, b := edge[0]-1, edge[1]-1
		adj[a][b] = true
		adj[b][a] = true
		deg[a] += 1
		deg[b] += 1
	}

	result := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !adj[i][j] {
				continue
			}
			for k := j + 1; k < n; k++ {
				if adj[i][k] && adj[k][j] {
					v := deg[i] + deg[j] + deg[k] - 6
					if result == -1 || v < result {
						result = v
					}
				}
			}
		}
	}

	return result
}
