package main

func originalDigits(s string) string {
	// zero, one, two, three, four, five, six, seven, eight, nine
	// first count the number of each character
	// some characters appear only in one number, so we know that number
	// must appear
	// by repeating this we can eliminate more and more characters
	// if there are still z -> add zero
	// if there are still x -> add six
	// if there are still w -> add two
	// if there are still s -> add seven
	// if there are still v -> add five
	// if there are still f -> add four
	// if there are still o -> add one
	// if there are still g -> add eight
	// if there are still n -> add nine
	// add three
	order := []string{"zero", "six", "two", "seven", "five", "four", "one", "eight", "nine", "three"}
	digit := []int{0, 6, 2, 7, 5, 4, 1, 8, 9, 3}
	charCount := make([]int, 26)
	for i := 0; i < len(s); i++ {
		charCount[s[i]-97] += 1
	}

	resultDigitCount := make([]int, 10)
	for i, num := range order {
		min := -1
		nc := map[byte]int{}
		for j := 0; j < len(num); j++ {
			nc[num[j]-97] += 1
		}
		for c, v := range nc {
			x := charCount[c] / v
			if min == -1 || x < min {
				min = x
			}
		}
		if min > 0 {
			resultDigitCount[digit[i]] = min
			for c, v := range nc {
				charCount[c] -= min * v
			}
		}
	}

	result := make([]byte, 0)
	for i, v := range resultDigitCount {
		for j := 0; j < v; j++ {
			result = append(result, byte(i)+byte('0'))
		}
	}
	return string(result)
}
