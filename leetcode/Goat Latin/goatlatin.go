// Package goatlatin solve leetcode challenge
// https://leetcode.com/contest/weekly-contest-82/problems/goat-latin/
package goatlatin

import (
	"fmt"
	"strings"
)

func toGoatLatin(S string) string {
	var (
		result []string
		x      byte
	)
	vowel := []rune{'a', 'e', 'i', 'o', 'u'}
	data := strings.Split(S, " ")

MAIN:
	for i := range data {
		for _, r := range vowel {
			if strings.EqualFold(string(data[i][0]), string(r)) {
				result = append(result, fmt.Sprintf("%v%v%v", data[i], "ma", strings.Repeat("a", i+1)))
				continue MAIN
			}
		}

		x, data[i] = data[i][0], data[i][1:]
		result = append(result, fmt.Sprintf("%v%v%v%v", data[i], string(x), "ma", strings.Repeat("a", i+1)))
	}

	return strings.Join(result, " ")
}
