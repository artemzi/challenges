// https://leetcode.com/problems/find-peak-element/description/
package main

import "fmt"

/**
A peak element is an element that is greater than its neighbors.

	Given an input array nums, where nums[i] ≠ nums[i+1], find a peak element and return its index.

	The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.

	You may imagine that nums[-1] = nums[n] = -∞.

	Example 1:

		Input: nums = [1,2,3,1]
		Output: 2
		Explanation: 3 is a peak element and your function should return the index number 2.
		Example 2:

		Input: nums = [1,2,1,3,5,6,4]
		Output: 1 or 5
		Explanation: Your function can return either index number 1 where the peak element is 2,
					or index number 5 where the peak element is 6.
	Note:

	Your solution should be in logarithmic complexity.
*/

func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	var l int
	r := len(nums) - 1

	for l+1 < r { // sequence contain at least 3 elements
		// get middle
		m := l + (r-l)/2
		if nums[m-1] < nums[m] {
			if nums[m] < nums[m+1] {
				l = m
			} else {
				return m
			}
		} else {
			if nums[m] < nums[m+1] {
				l = m
			} else {
				r = m
			}
		}
	}

	if nums[l] > nums[r] {
		return l
	}

	return r
}

func main() {
	fmt.Printf("expected: %v\ngot: %v\n", 2, findPeakElement([]int{1, 2, 3, 1}))
	fmt.Printf("expected: %v\ngot: %v\n", 0, findPeakElement([]int{3, 2, 1}))
}
