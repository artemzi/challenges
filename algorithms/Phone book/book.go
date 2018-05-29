package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
Телефонная книга
Реализовать структуру данных, эффективно обрабатывающую запро-
сы вида add number name, del number и find number.
	Вход. Последовательность запросов вида add number
		name, del number и find number, где number — те-
		лефонный номер, содержащий не более семи знаков,
		а name — короткая строка.

	Выход. Для каждого запроса find number выведите соот-
		ветствующее имя или сообщите, что такой записи нет.

	Формат входа. Первая строка содержит число запросов n. Каждая из
		следующих n строк задаёт запрос в одном из трёх описанных вы-
		ше форматов.

	Формат выхода. Для каждого запроса find выведите в отдельной
		строке либо имя, либо «not found».

	Ограничения. 1 ≤ n ≤ 10 5 . Телефонные номера содержат не более
		семи цифр и не содержат ведущих нулей. Имена содержат только
		буквы латинского алфавита, не являются пустыми строками и
		имеют длину не больше 15. Гарантируется, что среди имён не
		встречается строка «not found».
*/

func main() {
	var (
		n    int
		data map[int]string
	)
	fmt.Scan(&n)

	input := make([]string, 0, n)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if "" == scanner.Text() {
			break
		}
		input = append(input, scanner.Text())
	}

	data = make(map[int]string, n)
	for _, val := range input {
		t := strings.Split(val, " ")
		phone, _ := strconv.Atoi(t[1])
		switch t[0] {
		case "add":
			data[phone] = t[2]
		case "find":
			if v, ok := data[phone]; ok {
				fmt.Println(v)
				break
			}
			fmt.Println("not found")
		case "del":
			data[phone] = "not found"
		}
	}
}
