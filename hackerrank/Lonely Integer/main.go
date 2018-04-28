// https://www.hackerrank.com/challenges/lonely-integer/problem
package main

import "fmt"

func lonelyInteger(d []string) (result string) {
	cache := make(map[string]int, len(d))
	for _, num := range d {
		cache[num]++
	}

	for i, res := range cache {
		if res == 1 {
			result = i
			return
		}
	}
	return
}

func main() {
	var n int
	fmt.Scan(&n)

	data := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	fmt.Printf("%v\n", lonelyInteger(data))
}
