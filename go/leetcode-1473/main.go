package main

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	// <= 100 houses
	// <= 20 colors
	// d(i, j, k) = min cost to get j neighborhoods for houses 0...i where last neighborhood has color k
	// or -1 if this is not possible
	// 100*100*20 and for each element <= 20
	// can do this top down
	mem := make(map[E]int)
	min := -1
	for i := 0; i < n; i++ {
		x := rec(m-1, target, i, mem, houses, cost, n)
		if x != -1 {
			if min == -1 || x < min {
				min = x
			}
		}
	}
	return min
}

// min cost to get j neighborhoods for houses 0....i where last hast color k
// or -1 if not possible
func rec(i, j, k int, mem map[E]int, houses []int, cost [][]int, n int) int {
	if i+1 < j {
		return -1
	}
	if i == 0 {
		if houses[i] == 0 {
			return cost[0][k]
		} else if houses[i] == k+1 {
			return 0
		} else {
			return -1
		}
	}
	if j == 1 {
		v := 0
		for z := 0; z <= i; z++ {
			if houses[z] == 0 {
				v += cost[z][k]
			} else if houses[z] != k+1 {
				v = -1
				break
			}
		}
		mem[E{i, j, k}] = v
		return v
	}

	if houses[i] != 0 && houses[i] != k+1 {
		return -1
	}

	if v, ok := mem[E{i, j, k}]; ok {
		return v
	}

	result := -1
	base := 0
	if houses[i] == 0 {
		base = cost[i][k]
	}
	x := rec(i-1, j, k, mem, houses, cost, n)
	if x != -1 {
		v := x + base
		if result == -1 || v < result {
			result = v
		}
	}
	// use j-1 neighborhoods before, where last has a color != c
	for z := 0; z < n; z++ {
		if z != k {
			x = rec(i-1, j-1, z, mem, houses, cost, n)
			if x != -1 {
				v := x + base
				if result == -1 || v < result {
					result = v
				}
			}
		}
	}

	mem[E{i, j, k}] = result
	return result
}

type E struct {
	i int
	j int
	k int
}
