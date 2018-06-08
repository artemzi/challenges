package main

import "fmt"

/**
Given two strings s and t, determine if they are isomorphic.

	Two strings are isomorphic if the characters in s can be replaced to get t.

	All occurrences of a character must be replaced with another character
	while preserving the order of characters. No two characters may map to the same
	character but a character may map to itself.
*/

func isIsomorphic(s string, t string) bool {
	if s == t {
		return true
	}

	// Your runtime beats 100.00 % of golang submissions
	ms := make([]byte, 256)
	mt := make([]byte, 256)
	for i := 0; i < len(s); i++ {
		if ms[s[i]] == 0 && mt[t[i]] == 0 {
			ms[s[i]] = t[i]
			mt[t[i]] = s[i]
		} else {
			if ms[s[i]] != t[i] || mt[t[i]] != s[i] {
				return false
			}
		}
	}
	// =========

	// Your runtime beats 75.76 % of golang submissions.
	// rs := []rune(s)
	// rt := []rune(t)

	// ms := make(map[rune]rune)
	// mt := make(map[rune]rune)
	// for i := 0; i < len(s); i++ {
	// 	if _, ok := ms[rs[i]]; !ok {
	// 		ms[rs[i]] = rt[i]
	// 	} else if ms[rs[i]] != rt[i] {
	// 		return false
	// 	}

	// 	if _, ok := mt[rt[i]]; !ok {
	// 		mt[rt[i]] = rs[i]
	// 	} else if mt[rt[i]] != rs[i] {
	// 		return false
	// 	}
	// }
	// ==========

	return true
}

func main() {
	fmt.Printf("Expected: %v, got: %v\n", true, isIsomorphic("egg", "add"))
	fmt.Printf("Expected: %v, got: %v\n", false, isIsomorphic("foo", "bar"))
	fmt.Printf("Expected: %v, got: %v\n", true, isIsomorphic("paper", "title"))
	fmt.Printf("Expected: %v, got: %v\n", false, isIsomorphic("ab", "aa"))
}
