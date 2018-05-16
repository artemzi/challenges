package main

import (
	"fmt"
)

/*
Высота дерева
Вычислить высоту данного дерева.
	Вход. 	Корневое дерево с вершинами {0, . . . , n−1}, заданное
			как последовательность parent0, . . . , parentn−1,
			где parenti —родитель i-й вершины.

	Выход.  Высота дерева.

======================

Формата входа. Первая строка содержит натуральное число n.
	Вторая строка содержит n целых неотрицательных чисел
	parent0, . . . , parentn−1. Для каждого 0 ≤ i ≤ n − 1, parenti —
	родитель вершины i; если parenti = −1, то i является корнем.
	Гарантируется, что корень ровно один. Гарантируется, что
	данная последовательность задаёт дерево.

Формат выхода. Высота дерева.

Ограничения. 1 ≤ n ≤ 105.
*/

var nodes map[int][]int

func main() {
	var (
		n int
	)
	fmt.Scan(&n)

	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	fmt.Printf("%v\n", getMaxDeph(data))
}

func getMaxDeph(data []int) int {
	buildNodes(data)
	return height(indexOf(-1, data), 1)
}

func height(root int, h int) int {
	max := h
	for _, c := range nodes[root] {
		tmp := height(c, h) + 1
		if max < tmp {
			max = tmp
		}
	}
	return max
}

func buildNodes(data []int) {
	nodes = make(map[int][]int, len(data))

	for child, e := range data {
		if e == -1 {
			continue
		}
		nodes[e] = append(nodes[e], child)
	}
}

func indexOf(e int, data []int) int {
	for k, v := range data {
		if e == v {
			return k
		}
	}
	return -1
}
