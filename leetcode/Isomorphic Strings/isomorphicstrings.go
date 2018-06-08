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
	return false
}

func main() {
	fmt.Printf("Expected: %v, got: %v\n", true, isIsomorphic("egg", "add"))
	fmt.Printf("Expected: %v, got: %v\n", false, isIsomorphic("foo", "bar"))
	fmt.Printf("Expected: %v, got: %v\n", true, isIsomorphic("paper", "title"))
}
