package main

import "fmt"

func main() {
	k := 2
	prices := []int{2, 4, 1}
	fmt.Println(maxProfit(k, prices))
	k = 2
	prices = []int{3, 2, 6, 5, 0, 3}
	fmt.Println(maxProfit(k, prices))
}

func maxProfit(k int, prices []int) int {
	// d(i, j) = most profit for prices 0....i using <= j transactions
	// d(i, j) = max of d(i-1, j) and prices[i]-prices[z] + d(z-1, j-1) (i.e. sell at i and buy at z)
	n := len(prices)
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, k+1)
	}
	for i := 1; i < n; i++ {
		pi := prices[i]
		for j := 1; j <= k; j++ {
			v := d[i-1][j]
			// sell at i
			for z := 0; z < i; z++ {
				if prices[z] < pi {
					g := pi - prices[z]
					if z > 0 {
						g += d[z-1][j-1]
					}
					if g > v {
						v = g
					}
				}
			}
			d[i][j] = v
		}
	}

	return d[n-1][k]
}
