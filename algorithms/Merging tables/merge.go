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

// Element represent disjoint set
type Element struct {
	Parent *Element
	Value  interface{}
}

// Makeset operation
func Makeset() *Element {
	e := new(Element)
	e.Parent = e
	return e
}

// Find operation
func Find(e *Element) *Element {
	if e.Parent == e {
		return e
	}
	e.Parent = Find(e.Parent)
	return e.Parent
}

// Union operation
func Union(e1, e2 *Element) {
	root1 := Find(e1)
	root2 := Find(e2)
	root1.Parent = root2
}

func main() {
	var (
		n, m, d, s int
	)
	fmt.Scan(&n)
	fmt.Scan(&m)
	sizes := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&sizes[i])
	}

	data := make([][]int, m) // [destination, source]
	for i := 0; i < m; i++ {
		fmt.Scan(&d)
		fmt.Scan(&s)
		data[i] = []int{d, s}
	}

	fmt.Printf("DUBUG:  %d %d\n%v\n%v\n", n, m, sizes, data)
}
