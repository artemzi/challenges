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

// Stack must have comments
type Stack struct {
	max []int
}

// Pop the top item of the stack and return it
func (s *Stack) Pop() {
	s.max = s.max[:len(s.max)-1]
}

// Push a value onto the top of the stack
func (s *Stack) Push(value int) {
	if 0 == len(s.max) || s.max[len(s.max)-1] < value {
		s.max = append(s.max, value)
	} else {
		s.max = append(s.max, s.max[len(s.max)-1])
	}
}

// Max return current maximum value in stack
// for now works only with `int` type
func (s *Stack) Max() int {
	return s.max[len(s.max)-1]
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

	s := &Stack{[]int{}}

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
