package main

import (
	"container/heap"
	"fmt"
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

// An Item is something we manage in a priority queue.
type Item struct {
	value    int64 // The value of the item; arbitrary.
	priority int64 // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int64, priority int64) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	var (
		n, m int
	)
	fmt.Scan(&n)
	fmt.Scan(&m)
	data := make([]int64, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&data[i])
	}

	proc := make(PriorityQueue, n)
	for i := 0; i < n; i++ {
		proc[i] = &Item{
			value:    int64(i),
			priority: 0,
			index:    i,
		}
	}
	heap.Init(&proc)

	for i := 0; i < m; i++ {
		p := heap.Pop(&proc).(*Item)
		fmt.Println(p.value, i)

		heap.Push(&proc, &Item{
			value:    p.value,
			priority: p.priority + data[i],
		})
	}
}
