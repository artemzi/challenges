// https://leetcode.com/contest/weekly-contest-81/problems/shortest-distance-to-a-character/
package main

import "fmt"

func shortestToChar(S string, C byte) []int {
	r := make([]int, len(S))
	idx := len(S)
	for i, ch := range S {
		if ch == rune(C) {
			idx = -i
		}
		r[i] = i + idx
	}

	idx = 2 * len(S)
	for i := len(S) - 1; i >= 0; i-- {
		ch := S[i]
		if ch == C {
			idx = i
		}
		if idx-i < r[i] {
			r[i] = idx - i
		}
	}
	return r
}

func main() {
	fmt.Printf("%v\n", shortestToChar("loveleetcode", []byte("e")[0]))
}
