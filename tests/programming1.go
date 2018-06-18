package main

import (
	"strings"
)

func Reverse(s string) string {
	ss := strings.Split(s, " ")
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}

	return strings.Join(ss, " ")
}

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	//This code read the necessary data form the standard input
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		fmt.Println(Reverse(line))
// 	}
// }
