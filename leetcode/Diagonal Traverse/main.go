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

const (
	UPRIGHT = iota
	DOWNLEFT
)

func findDiagonalOrder(matrix [][]int) []int {
	// approach:
	// - use two states: UPRIGHT and DOWNLEFT, to represent the two iteration
	// modes
	// - with that established, it's as easy as checking the various edge
	// conditions that would cause us to switch the iteration mode
	// - note that the iteration direction checks prevent us from going outside
	// the bounds of the matrix, except when the iteration is done.
	// - test this on rectangular matrices as well as square matrices

	if len(matrix) == 0 {
		return []int{}
	}

	if len(matrix) == 1 {
		// A matrix with one row is already in the proper order:
		// Example: [1, 2, 3]
		return matrix[0]
	}

	result := []int{}
	mode := UPRIGHT
	row, col := 0, 0
	m := len(matrix)
	n := len(matrix[0])

	for row >= 0 && row < m && col >= 0 && col < n {
		result = append(result, matrix[row][col])

		if mode == UPRIGHT && (row == 0 || col+1 >= n) {
			if col+1 >= n {
				row++
			} else {
				col++
			}
			mode = DOWNLEFT
			continue
		}

		if mode == DOWNLEFT && (col == 0 || row+1 >= m) {
			if row+1 >= m {
				col++
			} else {
				row++
			}
			mode = UPRIGHT
			continue
		}

		// increment normally, since we're not at an edge
		if mode == UPRIGHT {
			row--
			col++
		} else {
			row++
			col--
		}
	}

	return result
}

// func findDiagonalOrder(matrix [][]int) []int {
// 	if len(matrix) == 0 || len(matrix[0]) == 0 {
// 		return []int{}
// 	}
// 	m, n := len(matrix), len(matrix[0])

// 	check := true
// 	getNext := func(i, j int) (int, int) {
// 		if check {
// 			i--
// 			j++
// 			if 0 <= i && j < n {
// 				return i, j
// 			}
// 			check = false

// 			if i < 0 && j < n {
// 				return 0, j
// 			}
// 			return i + 2, j - 1
// 		}

// 		i++
// 		j--
// 		if i < m && 0 <= j {
// 			return i, j
// 		}
// 		check = true

// 		if i < m && 0 > j {
// 			return i, 0
// 		}
// 		return i - 1, j + 2
// 	}

// 	mn := m * n
// 	res := make([]int, mn)

// 	i, j := 0, 0
// 	for k := 0; k < mn; k++ {
// 		res[k] = matrix[i][j]
// 		i, j = getNext(i, j)
// 	}

// 	return res
// }

func main() {
	fmt.Printf("expected: %v\ngot: %v\n",
		[]int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		findDiagonalOrder([][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}))
}
