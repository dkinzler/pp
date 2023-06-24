package main

import "fmt"

func main() {
	rods := []int{1, 2, 3, 6}
	fmt.Println(tallestBillboard(rods))
	rods = []int{518, 99, 365, 338, 800, 869, 917, 386, 129, 382, 116}
	fmt.Println(tallestBillboard(rods))
}

func tallestBillboard(rods []int) int {
	// find disjoint subsets A and B of rods such that
	// c = sum(A) = sum(B) is maximal
	// n <= 20   <= 2^20 ~ 1M subsets
	// sum(rods) <= 5000
	// so result is <= 2500
	// suppose we want to check if there are two disjoint subsets
	// have sum k
	// dp
	// d(i, j, k) = true if it is possible to make disjoint sum j and k
	// using element 0,...,i
	// we can assume j <= k
	// d(0, 0, 0) = true
	// d(0, 0, rods[0]) = true
	// d(0, rods[0], 0) = true
	// d(i, j, k) = true if d(i-1, j, k) or d(i-1, j-rods[i], k) or d(i-1, k-rods[i], j)

	sum := 0
	for _, v := range rods {
		sum += v
	}
	max := sum / 2
	n := len(rods)

	d := make([][]bool, max+1)
	for i := 0; i <= max; i++ {
		d[i] = make([]bool, max+1)
	}
	q := []E{{0, 0}}
	d[0][0] = true
	if rods[0] <= max {
		d[0][rods[0]] = true
		q = append(q, E{0, rods[0]})
	}
	for i := 1; i < n; i++ {
		z := rods[i]
		ll := len(q)
		for qi := 0; qi < ll; qi++ {
			e := q[qi]
			if e.j+z <= max {
				a, b := e.j+z, e.k
				if a > b {
					a, b = b, a
				}
				if !d[a][b] {
					d[a][b] = true
					q = append(q, E{a, b})
				}
			}
			if e.k+z <= max {
				a, b := e.j, e.k+z
				if !d[a][b] {
					d[a][b] = true
					q = append(q, E{a, b})
				}
			}
		}
	}

	for i := max; i >= 0; i-- {
		if d[i][i] {
			return i
		}
	}
	return 0
}

type E struct {
	j int
	k int
}
