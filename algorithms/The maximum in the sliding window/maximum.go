package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	// Stack must have comment
	Stack struct {
		top    *node
		max    []int
		length int
	}
	node struct {
		value int
		prev  *node
	}
)

// New create a new stack
func New() *Stack {
	return &Stack{nil, nil, 0}
}

// Pop the top item of the stack and return it
func (s *Stack) Pop() {
	if s.length == 0 {
		os.Exit(0)
	}

	n := s.top
	s.top = n.prev
	s.max = s.max[:len(s.max)-1]
	s.length--
}

// Push a value onto the top of the stack
func (s *Stack) Push(value int) {
	n := &node{value, s.top}
	s.top = n

	if 0 == len(s.max) || s.max[len(s.max)-1] < value {
		s.max = append(s.max, value)
	} else {
		s.max = append(s.max, s.max[len(s.max)-1])
	}

	s.length++
}

// Max return current maximum value in stack
func (s *Stack) Max() int {
	return s.max[len(s.max)-1]
}

/**
Стек с поддержкой максимума
Реализовать стек с поддержкой операций push, pop и max.
	Вход. Последовательность запросов push, pop и max .

	Выход. Для каждого запроса max вывести максимальное
		число, находящее на стеке.

===========

Формата входа. Первая строка содержит число запросов q. Каждая
	из последующих q строк задаёт запрос в одном из следующих
	форматов: push v, pop, or max.

Формат выхода. Для каждого запроса max в

Примеры.
	Вход:
		8
		2 7 3 1 5 2 6 2
		4
	Выход:
		7 7 5 6 6

	Вход:
		3
		2 1 5
		1
	Выход:
		2 1 5

	Вход:
		3
		2 3 9
		3
	Выход:
		9
*/

func main() {
	var (
		n int
		m int
	)
	fmt.Scan(&n)
	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
	fmt.Scan(&m)

	fmt.Printf("%v\n", strings.Join(process(n, data, m), " "))
}

func process(n int, d []int, m int) (r []string) {
	r = make([]string, 0, n-m+1)

	for i := 0; i < n-m+1; i++ {
		s := New()
		for _, val := range d[i : m+i] {
			s.Push(val)
		}
		r = append(r, strconv.Itoa(s.Max()))
	}

	return
}
