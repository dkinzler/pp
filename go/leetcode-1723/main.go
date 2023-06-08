package main

import (
	"fmt"
	"os"
	"pp/io"
)

func main() {
	reader, err := io.NewLineReader("testcases.txt")
	defer reader.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	i := 0
	for {
		c, err := readTestCase(reader)
		if err != nil {
			if err != io.ErrorEOF {
				fmt.Println(err)
			}
			return
		}
		result := minimumTimeRequired(c.jobs, c.k)
		if result == c.expected {
			fmt.Println("case", i, "ok")
		} else {
			fmt.Printf("case %v incorrect, expected %v got %v", i, c.expected, result)
		}
		i += 1
	}
}

type testCase struct {
	jobs     []int
	k        int
	expected int
}

func readTestCase(reader *io.LineReader) (testCase, error) {
	var result testCase
	jobs, err := reader.ReadIntSlice()
	if err != nil {
		return result, err
	}
	k, err := reader.ReadInt()
	if err != nil {
		return result, err
	}
	expected, err := reader.ReadInt()
	if err != nil {
		return result, err
	}
	result.jobs = jobs
	result.k = k
	result.expected = expected
	return result, nil
}

func minimumTimeRequired(jobs []int, k int) int {
	// use DP, d(i,j) = min max worker time when using machines 0,...,i for the subset
	// of jobs encoded by bitmask j
	n := len(jobs)
	nn := 1 << n

	d := make([][]int, k)
	// for i=0, d[0][j] is the time it takes to run subset j on a single machine
	// i.e. just the sum of the times
	d[0] = make([]int, nn)
	for j := 1; j < nn; j++ {
		for z := 0; z < n; z++ {
			if (1<<z)&j != 0 {
				d[0][j] = jobs[z] + d[0][j-(1<<z)]
				break
			}
		}
	}

	for i := 1; i < k; i++ {
		d[i] = make([]int, nn)
		// d[i][0] is always 0, no jobs = no time taken
		for j := 1; j < nn; j++ {
			m := d[0][j]
			for z := 0; z <= j; z++ {
				if z&j == z {
					// worker i performs jobs z, rest perform j-z
					v := max(d[0][z], d[i-1][j-z])
					if v < m {
						m = v
					}
				}
			}
			d[i][j] = m
		}
	}

	return d[k-1][nn-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
