package main

import (
	"fmt"
)

/*
Ваша цель в данной задаче — реализовать симуляцию объединения
таблиц в базе данных.
	Формат входа. Первая строка содержит числа n и m — число таблиц
		и число запросов, соответственно. Вторая строка содержит n це-
		лых чисел r 1 , . . . , r n — размеры таблиц. Каждая из последующих
		m строк содержит два номера таблиц destination i и source i , кото-
		рые необходимо объединить.

	Формат выхода. Для каждого из m запросов выведите максималь-
		ный размер таблицы после соответствующего объединения.

	Ограничения. 1 ≤ n, m ≤ 100 000; 0 ≤ r i ≤ 10 000;
		1 ≤ destination i , source i ≤ n.

	Пример.
		Вход:
			5
			1
			3
			2
			1
			5
			5
			5
			1 1 1 1
			5
			4
			4
			4
			3
		Выход:
			2
			2
			3
			5
			5

*/

var (
	max    int
	parent []int
)

// Find operation
func Find(e int) int {
	if e != parent[e] {
		parent[e] = Find(parent[e])
	}
	return parent[e]
}

// Max must have comment
func Max(data []int) int {
	m := 0
	for i := range data {
		if m < data[i] {
			m = data[i]
		}
	}
	return m
}

// Union operation
func Union(dest int, source int, sizes []int) int {
	pd := Find(dest)
	ps := Find(source)
	if pd != ps {
		parent[ps] = parent[pd]
		sizes[pd] += sizes[ps]
	}
	if sizes[pd] > max {
		max = sizes[pd]
	}
	return max
}

func main() {
	var (
		n, m, d, s int
	)
	fmt.Scan(&n)
	fmt.Scan(&m)
	sizes := make([]int, n)
	parent = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&sizes[i])
		parent[i] = i
	}

	data := make([][]int, m) // [destination, source]
	for i := 0; i < m; i++ {
		fmt.Scan(&d)
		fmt.Scan(&s)
		data[i] = []int{d, s}
	}

	max = Max(sizes)
	for _, operation := range data {
		d, s = operation[0], operation[1]
		fmt.Printf("%d\n", Union(d-1, s-1, sizes))
	}
}
