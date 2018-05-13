// Package flipimage leetcode contest
// https://leetcode.com/contest/weekly-contest-84/problems/flipping-an-image/
package flipimage

// Image must have comment
type Image [][]int

func flipAndInvertImage(A [][]int) [][]int {
	img := append(Image(nil), A...)

	for i, val := range img {
		img[i] = reverse(val)
	}

	for _, val := range img {
		for i, d := range val {
			val[i] = invert(d)
		}
	}
	return img
}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func invert(d int) int {
	if d == 1 {
		return 0
	}

	return 1
}
