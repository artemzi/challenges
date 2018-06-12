package main

import (
	"fmt"
	"sort"
)

/**
In a given integer array nums, there is always exactly one largest element.
	Find whether the largest element in the array is at least twice as much
	as every other number in the array.

	If it is, return the index of the largest element, otherwise return -1.
*/

func dominantIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	m := make(map[int]int, len(nums)) // [element : index]
	for i, e := range nums {
		m[e] = i
	}
	sort.Ints(nums)
	if nums[len(nums)-1]/2 >= nums[len(nums)-2] {
		return m[nums[len(nums)-1]]
	}

	return -1
}

func main() {
	fmt.Printf("Expected: %v, got: %v\n", 1, dominantIndex([]int{3, 6, 1, 0}))
	fmt.Printf("Expected: %v, got: %v\n", -1, dominantIndex([]int{1, 2, 3, 4}))
	fmt.Printf("Expected: %v, got: %v\n", -1, dominantIndex([]int{1}))
}
