package main

import (
	"fmt"
)

/**
Обработка сетевых пакетов
Реализовать обработчик сетевых пакетов.
	Вход.
		Размер буфера size и число пакетов n, а так-
		же две последовательности arrival 1 , . . . , arrival n и
		duration 1 , . . . , duration n , обозначающих время поступ-
		ления и длительность обработки n пакетов.

	Выход.
		Для каждого из данных n пакетов необходимо
		вывести время начала его обработки или −1, если пакет
		не был обработан (это происходит в случае, когда пакет
		поступает в момент, когда в

===============

Формат входа.
	Первая строка входа содержит размера буфера size
	и число пакетов n. Каждая из следующих n строк содержит два
	числа: время arrival i прибытия i-го пакета и время duration i ,
	необходимое на его обработку. Гарантируется, что arrival 1 ≤
	arrival 2 ≤ · · · ≤ arrival n . При этом может оказаться, что
	arrival i−1 = arrival i . В таком случае считаем, что пакет i − 1 по-
	ступил раньше пакета i.

Формата выхода.
	Для каждого из n пакетов выведите время, когда
	процессор начал его обрабатывать, или −1, если пакет был от-
	брошен.

Ограничения.
	Все числа во входе целые. 1 ≤ size ≤ 10 5 ; 1 ≤ n ≤ 10 5 ;
	0 ≤ arrival i ≤ 1
*/

func main() {
	var (
		size     int
		count    int
		arrival  int
		duration int
		buffer   []int
		freeTime int
		timeExit int
	)
	fmt.Scan(&size, &count)

	buffer = make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&arrival, &duration)
		if len(buffer) > 0 {
			for arrival >= buffer[0] {
				freeTime = buffer[0]
				if len(buffer) > 1 {
					buffer = buffer[1:]
				} else {
					buffer = buffer[:0]
					break
				}
			}
			if len(buffer) < size {
				if len(buffer) > 0 {
					timeExit = max(arrival, buffer[len(buffer)-1])
				} else {
					timeExit = max(arrival, freeTime)
				}
				buffer = append(buffer, timeExit+duration)
				fmt.Println(timeExit)
			} else {
				fmt.Println(-1)
			}
		}
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
