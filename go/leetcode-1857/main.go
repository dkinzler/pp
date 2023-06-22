package main

func largestPathValue(colors string, edges [][]int) int {
	// use topological sort
	// if there is no topological sorting there is a cycle and return -1
	// use dynamic programming to keep track of the paths
	// for each node keep track of the max frequency of a path for each color
	// whenever we select a degree 0 node for the sorting
	// we iterate over its neighbors and update their value
	n := len(colors)
	k := 26

	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, k)
		d[i][toColorIndex(colors[i])] = 1
	}
	inDegree := make([]int, n)
	adj := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		inDegree[b] += 1
		adj[a] = append(adj[a], b)
	}

	q := make([]int, 0)
	for i, v := range inDegree {
		if v == 0 {
			q = append(q, i)
		}
	}

	result := 1
	seen := 0
	for len(q) > 0 {
		curr := q[0]
		seen += 1
		q = q[1:]
		for _, nb := range adj[curr] {
			inDegree[nb] -= 1
			if inDegree[nb] == 0 {
				q = append(q, nb)
			}
			// update dp values
			a, b := d[curr], d[nb]
			c := toColorIndex(colors[nb])
			b[c] = max(b[c], a[c]+1)
			for j := 0; j < k; j++ {
				b[j] = max(b[j], a[j])
				result = max(result, b[j])
			}
		}
	}

	if seen != n {
		return -1
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func toColorIndex(b byte) int {
	return int(b - 97)
}
