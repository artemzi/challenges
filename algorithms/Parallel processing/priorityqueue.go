package main

import (
	"fmt"
	"math"
)

/**
Параллельная обработка
	По данным n процессорам и m задач определите, для каждой из задач,
	каким процессором она будет обработана.
		Вход. Число процессоров n и последовательность чисел
			t 0 , . . . , t m−1 , где t i — время, необходимое на обработку i-й
			задачи.

		Выход. Для каждой задачи определите, какой процессор
			и в какое время начнёт её обрабатывать, предполагая, что
			каждая задача поступает на обработку первому освободив-
			шемуся процессору.

			===========

	Формат входа. Первая строка входа содержит числа n и m. Вторая
		содержит числа t 0 , . . . , t m−1 , где t i — время, необходимое на об-
		работку i-й задачи. Считаем, что и процессоры, и задачи нуме-
		руются с нуля.

	Формат выхода. Выход должен содержать ровно m строк: i-я (счи-
		тая с нуля) строка должна содержать номер процесса, который
		получит i-ю задачу на обработку, и время, когда это произойдёт.

	Ограничения. 1 ≤ n ≤ 10 5 ; 1 ≤ m ≤ 10 5 ; 0 ≤ t i ≤ 10 9 .
*/

func main() {
	var (
		n, m int64
	)
	fmt.Scan(&n)
	fmt.Scan(&m)
	tasks := make([]int64, m)
	for i := 0; int64(i) < m; i++ {
		fmt.Scan(&tasks[i])
	}

	p := make([]int64, 0, n) // processors
	for i := 0; i < int(n); i++ {
		p = append(p, int64(i))
	}

	for i := 0; i < int(m); i++ {
		fmt.Println(p[0]%n, math.Floor(float64(p[0]/n)))
		p[0] += tasks[i] * n
		j, l, r := 0, 1, 2
		for int64(l) < n && p[j] > p[l] || int64(r) < n && p[j] > p[r] {
			if int64(r) >= n || p[j] > p[l] && p[r] > p[l] {
				p[j], p[l] = p[l], p[j]
				j = l
			} else {
				p[j], p[r] = p[r], p[j]
				j = r
			}
			l = j*2 + 1
			r = l + 1
		}
	}
}

// Sample Input:

// 2 5
// 1 2 3 4 5
// Sample Output:

// 0 0
// 1 0
// 0 1
// 1 2
// 0 4

// Sample Input:

// 2 15
// 0 0 1 0 0 0 2 1 2 3 0 0 0 2 1
// Sample Output:

// 0 0
// 0 0
// 0 0
// 1 0
// 1 0
// 1 0
// 1 0
// 0 1
// 0 2
// 1 2
// 0 4
// 0 4
// 0 4
// 0 4
// 1 5
