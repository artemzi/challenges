package main

import (
	"fmt"
)

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

	for _, val := range process(n, data, m) {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

func process(n int, d []int, m int) (r []int) {
	in := make([]int, 0, m)
	out := make([]int, 0, m)
	r = make([]int, 0, m)

	for i := 0; i < m; i++ {
		in = append(in, d[0])
		d = d[1:]
	}

	for _, val := range in {
		out = append(out, val)
	}
	in = in[:0]
	r = append(r, max(out))

	for _, val := range d {
		in = append(in, val)

		if len(out) == 0 {
			for len(in) > 0 {
				out = append(out, in[0])
				in = in[1:]
			}
		}
		out = out[1:]

		mo := max(out)
		mi := max(in)
		if mo > mi {
			r = append(r, mo)
		} else {
			r = append(r, mi)
		}
		d = d[1:]
	}

	return
}

func max(d []int) int {
	m := 0
	for i := range d {
		if d[i] > m {
			m = d[i]
		}
	}
	return m
}
