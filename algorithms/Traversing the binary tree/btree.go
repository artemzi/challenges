package main

import "fmt"

/**
Обход двоичного дерева
	Построить in-order, pre-order и post-order обходы данного двоичного дерева.
		Вход. Двоичное дерево.

		Выход. Все его вершины в трёх разных порядках: in-order,
			pre-order и post-order.

		Формат входа. Первая строка содержит число вершин n. Вершины
			дерева пронумерованы числами от 0 до n−1. Вершина 0 является
			корнем. Каждая из следующих n строк содержит информацию о
			вершинах 0, 1, . . . , n − 1: i-я строка задаёт числа key i , left i и right i ,
			где key i — ключ вершины i, left i — индекс левого сына верши-
			ны i, а right i — индекс правого сына вершины i. Если у верши-
			ны i нет одного или обоих сыновей, соответствующее значение
			равно −1.

		Формат выхода. Три строки: in-order, pre-order и post-order обходы.

		Ограничения. 1 ≤ n ≤ 10 5 ; 0 ≤ key i ≤ 10 9 ; −1 ≤ left i , right i ≤ n − 1.
			Гарантируется, что вход задаёт корректное двоичное дерево: в
			частности, если left i 6 = −1 и right i 6 = −1, то left i 6 = right i ; ника-
			кая вершина не является сыном двух вершин; каждая вершина
			является потомком корня.

		Пример:
			вход:
				5
				4 1 2
				2 3 4
				5 -1 -1
				1 -1 -1
				3 -1 -1
			выход:
				1 2 3 4 5
				4 2 1 3 5
				1 3 2 5 4
*/

// Leaf must have comment
type Leaf struct {
	key   uint
	left  int
	right int
}

func preOrder(v []Leaf, i int) {
	fmt.Printf("%d ", v[i].key)
	if v[i].left != -1 {
		preOrder(v, v[i].left)
	}
	if v[i].right != -1 {
		preOrder(v, v[i].right)
	}
}

func inOrder(v []Leaf, i int) {
	if v[i].left != -1 {
		inOrder(v, v[i].left)
	}
	fmt.Printf("%d ", v[i].key)
	if v[i].right != -1 {
		inOrder(v, v[i].right)
	}
}

func postOrder(v []Leaf, i int) {
	if v[i].left != -1 {
		postOrder(v, v[i].left)
	}
	if v[i].right != -1 {
		postOrder(v, v[i].right)
	}
	fmt.Printf("%d ", v[i].key)
}

func main() {
	var (
		n, left, right int
		key            uint
		tree           []Leaf
	)
	fmt.Scan(&n)
	tree = make([]Leaf, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&key)
		fmt.Scan(&left)
		fmt.Scan(&right)
		tree[i] = Leaf{key, left, right}
	}

	inOrder(tree, 0)
	fmt.Println()
	preOrder(tree, 0)
	fmt.Println()
	postOrder(tree, 0)
	fmt.Println()
}
