package main

import "fmt"

func main() {
	t1 := []int{1, 3, 1, 2}
	fmt.Println(totalStrength(t1))
	t2 := []int{5, 4, 6}
	fmt.Println(totalStrength(t2))
	t3 := []int{13, 13, 12, 12, 13, 12}
	fmt.Println(totalStrength(t3))
	t4 := []int{14, 9, 14, 3}
	fmt.Println(totalStrength(t4))
}

func totalStrength(strength []int) int {
	mod := 1000000007

	n := len(strength)
	suffix := make([]int, n)
	suffix[n-1] = strength[n-1] % mod
	for i := n - 2; i >= 0; i-- {
		suffix[i] = (strength[i] + suffix[i+1]) % mod
	}

	stack := make([]E, 0)
	result := 0
	cm := 0
	totalSuffix := 0
	for i, v := range strength {
		count := 1
		s := suffix[i]
		for len(stack) > 0 && stack[len(stack)-1].v >= v {
			last := stack[len(stack)-1]
			count += last.c
			s = (s + last.s) % mod
			totalSuffix = (totalSuffix - last.s*last.v) % mod
			totalSuffix = (totalSuffix + mod) % mod
			cm = (cm - last.c*last.v) % mod
			cm = (cm + mod) % mod
			stack = stack[0 : len(stack)-1]
		}
		cm = (cm + count*v) % mod
		totalSuffix = (totalSuffix + s*v) % mod

		g := 0
		if i < n-1 {
			g = (totalSuffix - suffix[i+1]*cm) % mod
			g = (g + mod) % mod
		} else {
			g = totalSuffix
		}
		result = (result + g) % mod
		stack = append(stack, E{v: v, c: count, s: s})
	}

	return result
}

type E struct {
	v int
	c int
	s int
}
