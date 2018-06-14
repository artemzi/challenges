package main

import "fmt"

/**
Given a matrix of M x N elements (M rows, N columns),
	return all elements of the matrix in diagonal order as shown in the below image.

	Example:
	Input:
		[
			[ 1, 2, 3 ],
			[ 4, 5, 6 ],
			[ 7, 8, 9 ]
		]
	Output:  [1,2,4,7,5,3,6,8,9]

	Note:
		The total number of elements of the given matrix will not exceed 10,000.
*/

func findDiagonalOrder(matrix [][]int) []int {
	return []int{}
}

func main() {
	fmt.Printf("expected: %v, got: %v\n",
		[]int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		findDiagonalOrder([][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}))
}
