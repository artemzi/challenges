package main

import "fmt"

/**
Given a string, find the length of the longest substring without repeating characters.

	Examples:

		Given "abcabcbb", the answer is "abc", which the length is 3.

		Given "bbbbb", the answer is "b", with the length of 1.

		Given "pwwkew", the answer is "wke", with the length of 3.
		Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]int, len(s))
	data := []byte(s)
	var leftIndex, ret int
	for i := 0; i < len(s); i++ {
		if _, ok := set[data[i]]; ok {
			leftIndex = max(leftIndex, set[s[i]]+1)
		}
		set[data[i]] = i
		ret = max(ret, i-leftIndex+1)
	}
	return ret
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Printf("%v\n", lengthOfLongestSubstring(""))
	fmt.Printf("%v\n", lengthOfLongestSubstring("pwwkew"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("c"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("au"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("dvdf"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("rdovre"))
}

//  Easy to understand O(n) solution
