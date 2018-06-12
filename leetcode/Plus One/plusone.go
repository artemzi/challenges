package main

import (
	"fmt"
)

/**
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

	The digits are stored such that the most significant digit is at the head of the list,
	and each element in the array contain a single digit.

	You may assume the integer does not contain any leading zero, except the number 0 itself.
*/

// Your runtime beats 100.00 % of golang submissions.
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] > 9 && i > 0 {
			digits[i] = 0
			continue
		} else if digits[i] > 9 && i == 0 {
			digits[i] = 0
			digits = append([]int{1}, digits...)
			continue
		}
		return digits
	}
	return digits
}

func main() {
	fmt.Printf("Expected: %v, got: %v\n", []int{1, 2, 4}, plusOne([]int{1, 2, 3}))
	fmt.Printf("Expected: %v, got: %v\n", []int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))
	fmt.Printf("Expected: %v, got: %v\n", []int{1, 0}, plusOne([]int{9}))
	fmt.Printf("Expected: %v, got: %v\n", []int{1}, plusOne([]int{0}))
	fmt.Printf("Expected: %v, got: %v\n", []int{1, 0, 0}, plusOne([]int{9, 9}))
}
