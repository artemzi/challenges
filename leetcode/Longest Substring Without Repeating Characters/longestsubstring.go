package main

/**
Given a string, find the length of the longest substring without repeating characters.

	Examples:

		Given "abcabcbb", the answer is "abc", which the length is 3.

		Given "bbbbb", the answer is "b", with the length of 1.

		Given "pwwkew", the answer is "wke", with the length of 3.
		Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := 1
	data := []rune(s)

	for i := len(data); i > 0; i-- {
		for j := 0; j < i; j++ {
			if !findDuplicates(data[j:i]) {
				if max < len(data[j:i]) {
					max = len(data[j:i])
				}
			}
		}
	}
	return max
}

func findDuplicates(r []rune) bool {
	set := make(map[rune]bool)
	for _, val := range r {
		if _, ok := set[val]; ok {
			return true
		}
		set[val] = true
	}
	return false
}

func main() {
	// fmt.Printf("%v\n", lengthOfLongestSubstring(""))
	// fmt.Printf("%v\n", lengthOfLongestSubstring("pwwkew"))
	// fmt.Printf("%v\n", lengthOfLongestSubstring("c"))
	// fmt.Printf("%v\n", lengthOfLongestSubstring("au"))
	// fmt.Printf("%v\n", lengthOfLongestSubstring("dvdf"))
	// fmt.Printf("%v\n", lengthOfLongestSubstring("rdovre"))
}

//  Time Limit Exceeded on ~2^20
