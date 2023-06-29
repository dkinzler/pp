package main

func minScoreTriangulation(values []int) int {
	// n-sided polygon, each vertex has a value
	// cut into n-2 triangles
	// minimize sum of product of vertices of each triangle
	// 3 <= n <= 50
	// can do this using dp
	// d(i, j) = min sum when using vertices i...j
	n := len(values)
	mem := make([][]int, n)
	for i := 0; i < n; i++ {
		mem[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mem[i][j] = -1
		}
	}
	return dp(0, n-1, values, mem)
}

func dp(i, j int, values []int, mem [][]int) int {
	if j-i+1 < 3 {
		return 0
	}
	if mem[i][j] != -1 {
		return mem[i][j]
	}
	result := 1 << 62
	for z := i + 1; z < j; z++ {
		v := values[i]*values[j]*values[z] + dp(i, z, values, mem) + dp(z, j, values, mem)
		if v < result {
			result = v
		}
	}
	mem[i][j] = result
	return result
}
