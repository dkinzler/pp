package main

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	// just simulate the opening of boxes using a q
	const open, closed int = 1, 0

	n := len(status)
	hasKey := make([]bool, n)
	hasBox := make([]bool, n)
	used := make([]bool, n)

	q := make([]int, 0)
	for _, box := range initialBoxes {
		hasBox[box] = true
		if status[box] == open {
			q = append(q, box)
			used[box] = true
		}
	}

	result := 0
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		result += candies[curr]
		for _, box := range keys[curr] {
			hasKey[box] = true
			if hasBox[box] && !used[box] {
				q = append(q, box)
				used[box] = true
			}
		}
		for _, box := range containedBoxes[curr] {
			hasBox[box] = true
			if status[box] == open || hasKey[box] {
				if !used[box] {
					q = append(q, box)
					used[box] = true
				}
			}
		}
	}
	return result
}
