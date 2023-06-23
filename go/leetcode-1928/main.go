package main

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	// d(i, j) = min cost to get from 0 to i in <= j time

	n := len(passingFees)
	// costs are always >= 1, so we can assume value 0 means not set
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, maxTime+1)
	}
	d[0][0] = passingFees[0]
	for t := 1; t <= maxTime; t++ {
		for _, edge := range edges {
			x, y, z := edge[0], edge[1], edge[2]
			// move from x to y
			if t-z >= 0 {
				if d[x][t-z] > 0 {
					v := d[x][t-z] + passingFees[y]
					old := d[y][t]
					if old == 0 || v < old {
						d[y][t] = v
					}
				}
				if d[y][t-z] > 0 {
					v := d[y][t-z] + passingFees[x]
					old := d[x][t]
					if old == 0 || v < old {
						d[x][t] = v
					}
				}
			}
		}
	}
	result := -1
	for _, v := range d[n-1] {
		if v != 0 {
			if result == -1 || v < result {
				result = v
			}
		}
	}
	return result
}
