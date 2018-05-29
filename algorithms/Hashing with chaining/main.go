package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
Ваша программа должна поддерживать следующие
типы запросов:
	• add string: добавить строку string в таблицу. Если такая
		строка уже есть, проигнорировать запрос;
	• del string: удалить строку string из таблицы. Если такой
		строки нет, проигнорировать запрос;
	• find string: вывести «yes» или «no» в зависимости от того,
		есть в таблице строка string или нет;
	• check i: вывести i-й список (используя пробел в качестве раз-
		делителя); если i-й список пуст, вывести пустую строку.
		При добавлении строки в цепочку, строка должна добавляться в на-
		чало цепочки.

	Формат входа. Первая строка размер хеш-таблицы m. Следующая
		строка содержит количество запросов n. Каждая из последую-
		щих n строк содержит запрос одного из перечисленных выше
		четырёх типов.

	Формат выхода. Для каждого из запросов типа find и check выве-
		дите результат в отдельной строке.

	Ограничения. 1 ≤ n ≤ 10 5 ; n 5 ≤ m ≤ n. Все строки имеют длину
		от одного до пятнадцати и содержат только буквы латинского
		алфавита.
*/

var m int64

func step(b int64, e int, p int64) int64 {
	if 0 == e {
		return int64(1)
	}

	r := b
	for i := 1; i < e; i++ {
		r = r % p * b % p
	}
	return r
}

func pHash(s string) int {
	var h int64
	x := int64(263)
	p := int64(1000000007)

	for i := 0; i < len(s); i++ {
		h += int64(s[i]) * step(x, i, p)
		h %= p
	}
	return int(h % m)
}

func main() {
	var (
		n int
	)
	fmt.Scan(&m)
	fmt.Scan(&n)

	input := make([]string, 0, n)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if "" == scanner.Text() {
			break
		}
		input = append(input, scanner.Text())
	}

	data := make(map[int][]string, m)
MAIN:
	for _, val := range input {
		t := strings.Split(val, " ")
		switch t[0] {
		case "add":
			key := pHash(t[1])
			if _, ok := data[key]; ok {
				for _, val := range data[key] {
					if val == t[1] {
						continue MAIN
					}
				}
			}
			data[key] = append(data[key], t[1])
		case "find":
			key := pHash(t[1])
			if _, ok := data[key]; ok {
				for _, val := range data[key] {
					if val == t[1] {
						fmt.Println("yes")
						continue MAIN
					}
				}
			}
			fmt.Println("no")
		case "check":
			var result string
			i, _ := strconv.Atoi(t[1])
			for j := len(data[i]) - 1; j >= 0; j-- {
				if data[i][j] == "" {
					continue
				}
				result += fmt.Sprintf("%v ", data[i][j])
			}
			fmt.Println(result)
		case "del":
			key := pHash(t[1])
			for i := len(data[key]) - 1; i >= 0; i-- {
				if data[key][i] == t[1] {
					data[key][i] = ""
				}
			}
		}
	}
}
