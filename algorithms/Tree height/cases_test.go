package treeheight

type tree struct {
	n   int
	val []int
}

type treeheightTest struct {
	input    *tree
	expected int
}

var stringTestCases = []treeheightTest{
	{
		input: &tree{
			5,
			[]int{4, -1, 4, 1, 1},
		},
		expected: 3,
	},
	{
		input: &tree{
			5,
			[]int{-1, 0, 4, 0, 3},
		},
		expected: 4,
	},
}
