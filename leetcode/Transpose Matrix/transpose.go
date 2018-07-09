package main

import "fmt"

func transpose(A [][]int) [][]int {
	out := make([][]int, len(A[0]))
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			out[j] = append(out[j], A[i][j])
		}
	}

	return out
}

func main() {
	fmt.Printf("expected: %v\ngot: %v\n", [][]int{
		{
			1, 4, 7,
		},
		{
			2, 5, 6,
		},
		{
			3, 6, 9,
		},
	}, transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))

	fmt.Printf("expected: %v\ngot: %v\n", [][]int{
		{
			1, 4,
		},
		{
			2, 5,
		},
		{
			3, 6,
		},
	}, transpose([][]int{{1, 2, 3}, {4, 5, 6}}))
}
