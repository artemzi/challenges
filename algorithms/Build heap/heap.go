package main

import (
	"fmt"
)

/*
Построение кучи
	Переставить элементы заданного массива чисел так, чтобы он удовле-
	творял свойству мин-кучи.

	Вход. Массив чисел A[0 . . . n − 1].

	Выход. Переставить элементы массива так, чтобы выпол-
		нялись неравенства A[i] ≤ A[2i + 1] и A[i] ≤ A[2i + 2] для
		всех i.

		=========

Формат входа. Первая строка содержит число n. Следующая строка
	задаёт массив чисел A[0], . . . , A[n − 1].

Формат выхода. Первая строка выхода должна содержать число об-
	менов m, которое должно удовлетворять неравенству 0 ≤ m ≤
	4n. Каждая из последующих m строк должна задавать обмен двух
	элементов массива A. Каждый обмен задаётся парой различных
	индексов 0 ≤ i 6 = j ≤ n − 1. После применения всех обменов в
	указанном порядке массив должен превратиться в мин-кучу, то
	есть для всех 0 ≤ i ≤ n − 1 должны выполняться следующие два
	условия:
	• если 2i + 1 ≤ n − 1, то A[i] < A[2i + 1].
	• если 2i + 2 ≤ n − 1, то A[i] < A[2i + 2].
Ограничения. 1 ≤ n ≤ 10, 5; 0 ≤ A[i]
*/

var swaps [][]int

func main() {
	var (
		n int
	)
	fmt.Scan(&n)
	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	swaps = make([][]int, 0, 4*n)

	for i := n - 1; i >= 0; i-- {
		shiftDown(i, data)
	}

	fmt.Println(len(swaps))
	for _, v := range swaps {
		fmt.Println(v[0], v[1])
	}
}

func shiftDown(i int, data []int) {
	l := i*2 + 1
	r := i*2 + 2
	if r < len(data) && data[r] < data[l] && data[i] > data[r] {
		swap(i, r, &data)
		shiftDown(r, data)
	}
	if l < len(data) && data[l] < data[i] {
		swap(i, l, &data)
		shiftDown(l, data)
	}
}

func swap(c, p int, data *[]int) {
	swaps = append(swaps, []int{c, p})

	(*data)[p], (*data)[c] = (*data)[c], (*data)[p]
}
