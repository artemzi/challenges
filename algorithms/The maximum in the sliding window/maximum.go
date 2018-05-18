package main

import (
	"fmt"
)

/**
Максимум в скользящем окне
	Найти максимум в каждом окне размера m данного массива чисел
	A[1 . . . n].
		Вход. Массив чисел A[1 . . . n] и число 1 ≤ m ≤ n.

		Выход. Максимум подмассива A[i . . . i + m − 1] для всех 1 ≤ i ≤ n − m + 1.

	===========
	Формат входа. Первая строка входа содержит число n, вторая — мас-
		сив A[1 . . . n], третья — число m.

	Формат выхода. n − m + 1 максимумов, разделённых пробелами.

	Ограниченяи.
		1 ≤ n ≤ 10 5 , 1 ≤ m ≤ n, 0 ≤ A[i] ≤ 10 5 для всех 1 ≤ i ≤ n.

	Примеры.
		Вход:
			8
			2 7 3 1 5 2 6 2
			4
		Выход:
			7 7 5 6 6

		Вход:
			3
			2 1 5
			1
		Выход:
			2 1 5

		Вход:
			3
			2 3 9
			3
		Выход:
			9
*/

func main() {
	var (
		n int
		m int
	)
	fmt.Scan(&n)
	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
	fmt.Scan(&m)

	for _, val := range process(n, data, m) {
		fmt.Print(fmt.Sprintf("%v ", val))
	}
	fmt.Println()
}

func process(n int, d []int, m int) (r []int) {
	var queue []int
	r = make([]int, 0, n-m+1)

	for i := 0; i < n-m+1; i++ {
		queue = d[i : m+i]
		for _, v := range queue {
			if i >= len(r) {
				r = append(r, v)
				continue
			}
			if r[i] < v {
				r[i] = v
			}
		}
	}
	return
}

// goos: linux
// goarch: amd64
// BenchmarkFlipAndInvertImage-2   	20000000	       108 ns/op
// PASS
