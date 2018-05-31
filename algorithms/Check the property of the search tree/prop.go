package main

import "fmt"

/**
Проверка свойства дерева поиска
	Проверить, является ли данное двоичное дерево деревом поиска.

	Вход. Двоичное дерево.

	Выход. Проверить, является ли оно корректным деревом
		поиска: верно ли, что для любой вершины дерева её ключ
		больше всех ключей в левом поддереве данной вершины и
		меньше всех ключей в правом поддереве.

	Формат входа. Первая строка содержит число вершин n. Вершины
		дерева пронумерованы числами от 0 до n−1. Вершина 0 является
		корнем. Каждая из следующих n строк содержит информацию о
		вершинах 0, 1, . . . , n − 1: i-я строка задаёт числа key i , left i и right i ,
		где key i — ключ вершины i, left i — индекс левого сына верши-
		ны i, а right i — индекс правого сына вершины i. Если у верши-
		ны i нет одного или обоих сыновей, соответствующее значение
		равно −1.

	Формат выхода. Выведите «CORRECT», если дерево является кор-
		ректным деревом поиска, и «INCORRECT» в противном случае.

	Ограничения. 0 ≤ n ≤ 10 5 ; −2 31 < key i < 2 31 − 1; −1 ≤ left i , right i ≤
		n − 1. Гарантируется, что вход задаёт корректное двоичное де-
		рево: в частности, если left i 6 = −1 и right i 6 = −1, то left i 6 = right i ;
		никакая вершина не является сыном двух вершин; каждая вер-
		шина является потомком корня.
*/

// Leaf must have comment
type Leaf struct {
	key   int
	left  int
	right int
}

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

func checkBst(t *[]Leaf, i int, mn int, mx int) bool {
	tree := (*t)
	if tree[i].key < mn || tree[i].key > mx {
		return false
	}
	if tree[i].left == -1 && tree[i].right == -1 {
		return true
	}

	if tree[i].key == tree[i].left {
		return false
	}

	if tree[i].left == -1 && tree[i].right != -1 {
		return checkBst(&tree, tree[i].right, tree[i].key+1, mx)
	}
	if tree[i].left != -1 && tree[i].right == -1 {
		return checkBst(&tree, tree[i].left, mn, tree[i].key-1)
	}

	return checkBst(&tree, tree[i].left, mn, tree[i].key-1) &&
		checkBst(&tree, tree[i].right, tree[i].key+1, mx)
}

func main() {
	var (
		n, key, left, right int
		tree                []Leaf
	)
	fmt.Scan(&n)
	if n == 0 {
		fmt.Println("CORRECT")
	} else {
		tree = make([]Leaf, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&key)
			fmt.Scan(&left)
			fmt.Scan(&right)
			tree[i] = Leaf{key, left, right}
		}

		if checkBst(&tree, 0, minInt, maxInt) {
			fmt.Println("CORRECT")
		} else {
			fmt.Println("INCORRECT")
		}
	}
}
