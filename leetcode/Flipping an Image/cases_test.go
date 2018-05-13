package flipimage

type flipimageTest struct {
	input    [][]int
	expected [][]int
}

var stringTestCases = []flipimageTest{
	{
		input: [][]int{
			[]int{1, 1, 0},
			[]int{1, 0, 1},
			[]int{0, 0, 0},
		},
		expected: [][]int{
			[]int{1, 0, 0},
			[]int{0, 1, 0},
			[]int{1, 1, 1},
		},
		// Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
		// Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]
	},
	{
		input: [][]int{
			[]int{1, 1, 0, 0},
			[]int{1, 0, 0, 1},
			[]int{0, 1, 1, 1},
			[]int{1, 0, 1, 0},
		},
		expected: [][]int{
			[]int{1, 1, 0, 0},
			[]int{0, 1, 1, 0},
			[]int{0, 0, 0, 1},
			[]int{1, 0, 1, 0},
		},
		// Explanation: First reverse each row: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
		// Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
	},
}
