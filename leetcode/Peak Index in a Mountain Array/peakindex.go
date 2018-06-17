package main

import "fmt"

/**
Let's call an array A a mountain if the following properties hold:

	A.length >= 3

	There exists some 0 < i < A.length - 1 such that
		A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]

	Given an array that is definitely a mountain, return any i such that
		A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].
*/

func peakIndexInMountainArray(A []int) int {
	if 0 == len(A) {
		return 0
	}

	for i, e := range A {
		if e > A[i+1] {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Printf("expected: %v\ngot: %v\n", 1, peakIndexInMountainArray([]int{0, 1, 0}))
}
