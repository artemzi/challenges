package main

import (
	"fmt"
)

/**
Given an array of integers nums,
 	write a method that returns the "pivot" index of this array.

	We define the pivot index as the index where the sum of the numbers
	to the left of the index is equal to the sum of the numbers to the right
	of the index.

	If no such index exists, we should return -1.
	If there are multiple pivot indexes, you should return the left-most
	pivot index.
*/

// Your runtime beats 100.00 % of golang submissions
func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var sum, lsum int

	// get sum of all elements
	for _, num := range nums {
		sum += num
	}

	// make step from `0` to `i`,
	// before 2*`lsum` will be equal to total `sum`
	for i, num := range nums {
		if 2*lsum+num == sum {
			return i
		}

		lsum += num
	}

	return -1
}

// Your runtime beats 12.00 % of golang submissions
// func pivotIndex(nums []int) int {
// 	for i := 0; i < len(nums); i++ {
// 		if sum(nums[:i]) == sum(nums[i+1:]) {
// 			return i
// 		}
// 	}

// 	return -1
// }

// func sum(n []int) (r int) {
// 	for i := range n {
// 		r += n[i]
// 	}
// 	return
// }

func main() {
	fmt.Printf("Expected: %v, got: %v\n", 3, pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Printf("Expected: %v, got: %v\n", -1, pivotIndex([]int{1, 2, 3}))
	fmt.Printf("Expected: %v, got: %v\n", 0, pivotIndex([]int{-1, -1, -1, 0, 1, 1}))
}
