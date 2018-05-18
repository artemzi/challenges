package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
Стек с поддержкой максимума
Реализовать стек с поддержкой операций push, pop и max.

	Вход. Последовательность запросов push, pop и max .

	Выход. Для каждого запроса max вывести максимальное
		число, находящее на стеке.

=========

Формата входа. Первая строка содержит число запросов q. Каждая
	из последующих q строк задаёт запрос в одном из следующих
	форматов: push v, pop, or max.

Формат выхода. Для каждого запроса max
*/

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
	s.max = append(s.max, n.value)
	s.length++
}

// Max return current maximum value in stack
// for now works only with `int` type
func (s *Stack) Max() int {
	var m int
	for _, v := range s.max {
		if v > m {
			m = v
		}
	}
	return m
}

func main() {
	var n int
	fmt.Scan(&n)

	data := make([]string, 0, n)
	input := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		input.Scan()
		data = append(data, input.Text())
	}

	s := New() // initialize stack

	for _, command := range data {
		c := strings.Split(command, " ")
		if c[0] == "push" {
			val, _ := strconv.Atoi(c[1])
			s.Push(val)
			continue
		}
		if c[0] == "pop" {
			s.Pop()
			continue
		}

		fmt.Printf("%v\n", s.Max())
	}
}
