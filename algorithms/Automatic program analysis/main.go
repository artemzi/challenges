package main

import (
	"fmt"
	"os"
)

/*
Система равенств и неравенств
	Проверить, можно ли присвоить переменным целые значения, чтобы
	выполнить заданные равенства вида x i = x j и неравенства вида x p 6 = x q .

	Вход. Число переменных n, а также список равенств вида
		x i = x j и неравенства вида x p 6 = x q .

	Выход. Проверить, выполнима ли данная система.

	Формат входа. Первая строка содержит числа n, e, d. Каждая из сле-
		дующих e строк содержит два числа i и j и задаёт равенство
		x i = x j . Каждая из следующих d строк содержит два числа i и j и
		задаёт неравенство x i 6 = x j . Переменные индексируются с 1:
		x 1 , . . . , x n .

	Формат выхода. Выведите 1, если переменным x 1 , . . . , x n можно
		присвоить целые значения, чтобы все равенства и неравенства
		выполнились. В противном случае выведите 0.

	Ограничения. 1 ≤ n ≤ 10 5 ; 0 ≤ e, d; e + d ≤ 2 · 10 5 ; 1 ≤ i, j ≤ n.
*/

var (
	a, rank []int
)

func union(l int, r int) {
	lp := find(l)
	rp := find(r)
	if rank[lp] > rank[rp] {
		a[rp] = lp
	} else {
		a[lp] = rp
		rank[lp]++
	}
}

func find(e int) int {
	p := e
	for {
		if p == a[p] {
			break
		}
		p = a[p]
	}
	return a[p]
}

func main() {
	var (
		n, e, d, inp1, inp2 int
	)
	fmt.Scan(&n)
	fmt.Scan(&e)
	fmt.Scan(&d)

	a = make([]int, n)
	rank = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
		rank[i] = 0
	}

	for i := 0; i < e; i++ {
		fmt.Scan(&inp1)
		fmt.Scan(&inp2)

		inp1--
		inp2--
		union(inp1, inp2)
	}

	for i := 0; i < d; i++ {
		fmt.Scan(&inp1)
		fmt.Scan(&inp2)

		inp1--
		inp2--
		if find(inp1) == find(inp2) {
			fmt.Println(0)
			os.Exit(0)
		}
	}
	fmt.Println(1)
}
