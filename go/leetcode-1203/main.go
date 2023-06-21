package main

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	// this is a topological sort
	// but we want each group together
	// so we can first topologically sort the groups
	// and then sort each group

	// for each node without a group create a new one
	groupSize := make(map[int]int)
	for x, g := range group {
		if g == -1 {
			group[x] = m
			m += 1
		}
		groupSize[group[x]] += 1
	}

	// create a graph of groups
	inDegree := make([]int, m)
	graph := make([][]int, m)
	for i, b := range beforeItems {
		for _, j := range b {
			// node j before node i
			// i.e edge from j to i
			g1, g2 := group[i], group[j]
			if g1 != g2 {
				inDegree[g1] += 1
				graph[g2] = append(graph[g2], g1)
			}
		}
	}
	sortedGroups := topologicalSort(inDegree, graph)
	if sortedGroups == nil {
		return nil
	}

	// create a graph where there are only edges between groups
	inDegree = make([]int, n)
	graph = make([][]int, n)
	for i, b := range beforeItems {
		for _, j := range b {
			// node j before node i
			// i.e edge from j to i
			g1, g2 := group[i], group[j]
			if g1 == g2 {
				inDegree[i] += 1
				graph[j] = append(graph[j], i)
			}
		}
	}
	result := make([]int, 0)
	for _, g := range sortedGroups {
		r := topologicalSortGroup(inDegree, graph, group, g)
		if len(r) == groupSize[g] {
			result = append(result, r...)
		} else {
			return nil
		}
	}
	return result
}

func topologicalSort(inDegree []int, graph [][]int) []int {
	result := make([]int, 0)
	q := make([]int, 0)
	for node, deg := range inDegree {
		if deg == 0 {
			q = append(q, node)
		}
	}
	for len(q) > 0 {
		curr := q[0]
		result = append(result, curr)
		q = q[1:]
		for _, nb := range graph[curr] {
			inDegree[nb] -= 1
			if inDegree[nb] == 0 {
				q = append(q, nb)
			}
		}
	}

	if len(result) == len(inDegree) {
		return result
	}
	return nil
}

func topologicalSortGroup(inDegree []int, graph [][]int, group []int, target int) []int {
	result := make([]int, 0)
	q := make([]int, 0)
	for node, deg := range inDegree {
		if group[node] == target {
			if deg == 0 {
				q = append(q, node)
			}
		}
	}
	for len(q) > 0 {
		curr := q[0]
		result = append(result, curr)
		q = q[1:]
		for _, nb := range graph[curr] {
			inDegree[nb] -= 1
			if inDegree[nb] == 0 {
				q = append(q, nb)
			}
		}
	}

	return result
}
