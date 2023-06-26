package main

import "fmt"

func main() {
	stones := []int{0, 1, 3, 5, 6, 8, 12, 17}
	fmt.Println(canCross(stones))
	stones = []int{0, 1, 2, 3, 4, 8, 9, 11}
	fmt.Println(canCross(stones))
}

func canCross(stones []int) bool {
	// d(i, j) = true if we can make it to pos i via pos j, j < i
	// d(i, 0) = true if stones[i] == 1
	// for d(i, j) we would have to make a jump of length k=stones[i]-stones[j]
	// i.e. we would need a prev jump with length k-1, k or k+1
	// i.e. d(i, j) = true if d(j, z) where stones[j]-stones[z] = k-1/k/k+1
	// i.e. we need to find an appropriate z
	// all the values are different so we can just use a map of the ones we have already seen to find z
	n := len(stones)
	valueToIndex := make(map[int]int)
	for i, v := range stones {
		valueToIndex[v] = i
	}

	d := make([][]bool, n)
	for i := 0; i < n; i++ {
		d[i] = make([]bool, n)
	}
	if stones[1] == 1 {
		d[1][0] = true
	} else {
		return false
	}
	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			k := stones[i] - stones[j]
			for m := -1; m <= 1; m++ {
				z := stones[j] - k + m
				if zi, ok := valueToIndex[z]; ok {
					if zi < j && d[j][zi] {
						d[i][j] = true
						break
					}
				}
			}
		}
	}

	for _, v := range d[n-1] {
		if v {
			return true
		}
	}
	return false
}
