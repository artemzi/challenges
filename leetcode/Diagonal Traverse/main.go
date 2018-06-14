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
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])

	check := true
	getNext := func(i, j int) (int, int) {
		if check {
			i--
			j++
			if 0 <= i && j < n {
				return i, j
			}
			check = false

			if i < 0 && j < n {
				return 0, j
			}
			return i + 2, j - 1
		}

		i++
		j--
		if i < m && 0 <= j {
			return i, j
		}
		check = true

		if i < m && 0 > j {
			return i, 0
		}
		return i - 1, j + 2
	}

	mn := m * n
	res := make([]int, mn)

	i, j := 0, 0
	for k := 0; k < mn; k++ {
		res[k] = matrix[i][j]
		i, j = getNext(i, j)
	}

	return res
}

func main() {
	fmt.Printf("expected: %v\ngot: %v\n",
		[]int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		findDiagonalOrder([][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}))
}
