package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var data []int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 21)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, row := range lines {
		for _, v := range strings.Split(row, " ") {
			if v == "" {
				continue
			}
			i, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, i)
		}
	}

	var ans int
	for i, _ := range data {
		if v := highest(i); v > ans {
			ans = v
		}
	}
	fmt.Print(ans)
}

func directions(i int) [][]int {
	const row = 20
	return [][]int{
		{i, i - 1, i - 2, i - 3},
		{i, i + 1, i + 2, i + 3},
		{i, i - (row * 1), i - (row * 2), i - (row * 3)},
		{i, i + (row * 1), i + (row * 2), i + (row * 3)},
		{i, i - (row * 1) - 1, i - (row * 2) - 2, i - (row * 3) - 3},
		{i, i - (row * 1) + 1, i - (row * 2) + 2, i - (row * 3) + 3},
		{i, i + (row * 1) - 1, i + (row * 2) - 2, i + (row * 3) - 3},
		{i, i + (row * 1) + 1, i + (row * 2) + 2, i + (row * 3) + 3},
	}
}

func highest(i int) int {
	var largest int

Direction:
	for _, direction := range directions(i) {
		score := 1
		for _, offset := range direction {
			if i+offset < 0 || i+offset > len(data)-1 {
				continue Direction
			}
			score = score * data[i+offset]
		}
		if score > largest {
			largest = score
		}
	}
	return largest
}
