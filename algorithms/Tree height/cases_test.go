package main

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
	{
		input: &tree{
			10,
			[]int{9, 7, 5, 5, 2, 9, 9, 9, 2, -1},
		},
		expected: 4,
	},
}
