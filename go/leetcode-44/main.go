package main

import (
	"fmt"
	"strings"
)

func main() {
	s, p := "aa", "a"
	fmt.Println(isMatch(s, p))
	s, p = "aa", "*"
	fmt.Println(isMatch(s, p))
	s, p = "cb", "?a"
	fmt.Println(isMatch(s, p))
	s, p = "abcabczzzde", "*abc???de*"
	fmt.Println(isMatch(s, p))
	s, p = "abcd", "abc*d"
	fmt.Println(isMatch(s, p))
}

func isMatch(s string, p string) bool {
	// the idea is we have fixed sequence we need to match separated by *
	// also: multiple * in sequence are equivalent to a single *
	parts := make([]string, 0)
	for _, v := range strings.Split(p, "*") {
		if len(v) > 0 {
			parts = append(parts, v)
		}
	}
	if len(p) == 0 {
		return len(s) == 0
	}
	starCount := 0
	for i := 0; i < len(p); i++ {
		if p[i] == byte('*') {
			starCount += 1
		}
	}
	if starCount == len(p) {
		return true
	}
	if starCount == 0 {
		if len(s) != len(p) {
			return false
		}
		return matches(s, p, 0, len(s)-1)
	}

	fixedStart := false
	if p[0] != byte('*') {
		fixedStart = true
	}
	fixedEnd := false
	if p[len(p)-1] != byte('*') {
		fixedEnd = true
	}

	if fixedStart && len(parts) == 1 {
		pp := parts[0]
		if len(pp) > len(s) {
			return false
		}
		return matches(s, pp, 0, len(pp)-1)
	}
	if fixedEnd && len(parts) == 1 {
		pp := parts[0]
		if len(pp) > len(s) {
			return false
		}
		return matches(s, pp, len(s)-1-len(pp)+1, len(s)-1)
	}

	currStart := 0
	currEnd := len(s) - 1
	if fixedStart {
		pp := parts[0]
		if !matches(s, pp, currStart, currEnd) {
			return false
		}
		currStart = len(pp)
	}
	if fixedEnd {
		pp := parts[len(parts)-1]
		ee := currEnd - len(pp) + 1
		if ee < currStart {
			return false
		}
		if !matches(s, pp, ee, currEnd) {
			return false
		}
		currEnd = currEnd - len(pp)
	}

	pi := 0
	if fixedStart {
		pi += 1
	}
	pe := len(parts) - 1
	if fixedEnd {
		pe -= 1
	}

	for i := pi; i <= pe; i++ {
		pp := parts[i]
		if len(pp) == 0 {
			continue
		}
		j, ok := matchesAnywhere(s, pp, currStart, currEnd)
		if !ok {
			return false
		}
		currStart = j
	}
	return true
}

func matchesAnywhere(s string, p string, start, end int) (int, bool) {
	for i := start; i+len(p)-1 <= end; i++ {
		if matches(s, p, i, end) {
			return i + len(p), true
		}
	}
	return -1, false
}

// p should only contain ? and lowercase english letters
func matches(s string, p string, start, end int) bool {
	if start+len(p)-1 > end {
		return false
	}
	for i := 0; i < len(p); i++ {
		if p[i] == byte('?') {
			continue
		}
		if p[i] != s[start+i] {
			return false
		}
	}
	return true
}
