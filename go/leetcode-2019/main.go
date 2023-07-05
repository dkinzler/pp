package main

import "fmt"

func main() {
	s := "7+3*1*2"
	answers := []int{20, 13, 42}
	fmt.Println(scoreOfStudents(s, answers))
	s = "3+5*2"
	answers = []int{13, 0, 10, 13, 13, 16, 16}
	fmt.Println(scoreOfStudents(s, answers))
	s = "6+0*1"
	answers = []int{12, 9, 6, 4, 8, 6}
	fmt.Println(scoreOfStudents(s, answers))
	s = "3+6+3*9+3*9+9*6+9+6*9+3*6+3*6+9"
	answers = []int{12, 9, 6, 4, 8, 6}
	fmt.Println(scoreOfStudents(s, answers))
}

func scoreOfStudents(s string, answers []int) int {
	// note that s always has uneven length
	// operators are at odd indices
	// we basically need to compute which values are possible
	// using different evaluation orders
	// can do dp
	// d(i, j) = set of values possible for range i, j where both are even indices
	// for i = j this is just the value
	// otherwise we can choose each operator i+1, i+3, ... j-1 as the last
	n := len(s)/2 + 1

	d := make([][]map[int]bool, n)
	for i := 0; i < n; i++ {
		d[i] = make([]map[int]bool, n)
		for j := i; j < n; j++ {
			d[i][j] = map[int]bool{}
		}
	}
	for i := 0; i < n; i++ {
		d[i][i][int(s[2*i])-48] = true
	}
	for k := 2; k <= n; k++ {
		for i := 0; i+k-1 < n; i++ {
			j := i + k - 1
			// which ops are there between i and j
			// i,...,j-1
			// for an op index z, the corresponding pos in s is
			// 2*z+1, e.g. z=0 -> 1, z=1 -> 3 ...
			mm := d[i][j]
			for z := i; z < j; z++ {
				isMult := false
				if s[2*z+1] == byte('*') {
					isMult = true
				}
				left := d[i][z]
				right := d[z+1][j]
				for v1 := range left {
					for v2 := range right {
						var v int
						if isMult {
							v = v1 * v2
						} else {
							v = v1 + v2
						}
						// we know all the answers are <= 1000
						// so any value > 1000 cannot lead to a result
						// since multiplication and addition can only increase values
						// unless we multiply by 0, but then it doesn't matter what the number is
						// we can all treat them as 1001
						if v > 1000 {
							v = 1001
						}
						mm[v] = true
					}
				}
			}
		}
	}

	correct := eval(s)
	result := 0
	for _, a := range answers {
		if a == correct {
			result += 5
		} else if _, ok := d[0][n-1][a]; ok {
			result += 2
		}
	}

	return result
}

func eval(s string) int {
	stack := make([]int, 0)
	stack = append(stack, int(s[0])-48)
	for i := 2; i < len(s); i += 2 {
		isMult := false
		if s[i-1] == byte('*') {
			isMult = true
		}
		x := int(s[i]) - 48
		if isMult {
			last := len(stack) - 1
			stack[last] = stack[last] * x
		} else {
			stack = append(stack, x)
		}
	}
	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}
