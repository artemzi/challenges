package main

import (
	"fmt"
	"reflect"
)

/**
Given two strings A and B of lowercase letters,
	return true if and only if we can swap two letters in
	A so that the result equals B.

	Example 1:

		Input: A = "ab", B = "ba"
		Output: true
	Example 2:

		Input: A = "ab", B = "ab"
		Output: false
	Example 3:

		Input: A = "aa", B = "aa"
		Output: true
	Example 4:

		Input: A = "aaaaaaabc", B = "aaaaaaacb"
		Output: true
	Example 5:

		Input: A = "", B = "aa"
		Output: false
*/

func buddyStrings(A string, B string) bool {
	if reflect.DeepEqual(A, B) {
		c := make([]int, 26)
		a := []rune(A)
		for i := 0; i < len(A); i++ {
			c[a[i]-'a']++
		}

		for _, count := range c {
			if count > 1 {
				return true
			}
		}
		return false
	} else {
		f, s := -1, -1
		a, b := []rune(A), []rune(B)
		for i := 0; i < len(A); i++ {
			if a[i] != b[i] {
				if f == -1 {
					f = i
				} else if s == -1 {
					s = i
				} else {
					return false
				}
			}
		}

		return (s != -1 && a[f] == b[s] && a[s] == b[f])
	}
}

func main() {
	fmt.Printf("expected: %v, got: %v\n", true, buddyStrings("ab", "ba"))
	fmt.Printf("expected: %v, got: %v\n", false, buddyStrings("ab", "ab"))
	fmt.Printf("expected: %v, got: %v\n", true, buddyStrings("aa", "aa"))
	fmt.Printf("expected: %v, got: %v\n", true, buddyStrings("aaaaaaabc", "aaaaaaacb"))
	fmt.Printf("expected: %v, got: %v\n", false, buddyStrings("", "aa"))
}

// 20 / 20 test cases passed
