package main

import "fmt"

func matrixScore(A [][]int) int {
	var ans int
	R := len(A)
	C := len(A[0])

	for c := 0; c < C; c++ {
		col := 0
		for r := 0; r < R; r++ {
			col += A[r][c] ^ A[r][0]
		}
		ans += max(col, R-col) * int(1<<(uint(C-1-c)))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Printf("expected: %v\ngot: %v\n", 39, matrixScore([][]int{
		[]int{0, 0, 1, 1},
		[]int{1, 0, 1, 0},
		[]int{1, 1, 0, 0},
	}))
}
