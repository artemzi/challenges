package main

import "fmt"

func lemonadeChange(bills []int) bool {
	var f, t int
	for _, b := range bills {
		if b == 5 {
			f++
		} else if b == 10 {
			if f == 0 {
				return false
			}
			f--
			t++
		} else {
			if f > 0 && t > 0 {
				f--
				t--
			} else if f >= 3 {
				f -= 3
			} else {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Printf("expected: %v\ngot: %v\n", true, lemonadeChange([]int{5, 5, 5, 10, 20}))
	fmt.Printf("expected: %v\ngot: %v\n", true, lemonadeChange([]int{5, 5, 10}))
	fmt.Printf("expected: %v\ngot: %v\n", false, lemonadeChange([]int{10, 10}))
}
