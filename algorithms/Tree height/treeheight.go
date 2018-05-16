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

func main() {
	var n int
	fmt.Scan(&n)

	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	fmt.Printf("%v\n", getMaxDeph(data))
}

func getMaxDeph(data []int) int {
	var root int
	nodes := make(map[int][]int, len(data))

	for child, e := range data {
		if e == -1 {
			root = child
			continue
		}
		nodes[e] = append(nodes[e], child)
	}
	return height(root, 1, nodes)
}

func height(root int, h int, nodes map[int][]int) int {
	max := h
	for _, c := range nodes[root] {
		tmp := height(c, h, nodes) + 1
		if max < tmp {
			max = tmp
		}
	}
	return max
}
