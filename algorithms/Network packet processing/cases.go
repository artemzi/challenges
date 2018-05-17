package main

type processTest struct {
	skipped  bool
	input    [][]int
	expected []int
}

var stringTestCases = []processTest{
	{
		skipped: false,
		input: [][]int{
			[]int{1, 0},
		},
		expected: []int{},
	},
	{
		skipped: false,
		input: [][]int{
			[]int{1, 1},
			[]int{0, 0},
		},
		expected: []int{0},
	},
	{
		skipped: false,
		input: [][]int{
			[]int{1, 2},
			[]int{0, 1},
			[]int{0, 1},
		},
		expected: []int{
			0,
			-1,
		},
	},
	{
		skipped: false,
		input: [][]int{
			[]int{1, 2},
			[]int{0, 1},
			[]int{1, 1},
		},
		expected: []int{
			0,
			1,
		},
	},
	{
		skipped: false,
		input: [][]int{
			[]int{2, 3},
			[]int{1, 100},
			[]int{1, 100},
			[]int{1, 0},
		},
		expected: []int{
			1,
			100,
			-1,
		},
	},
	{
		skipped: false,
		input: [][]int{
			[]int{2, 6},
			[]int{0, 2},
			[]int{0, 0},
			[]int{2, 0},
			[]int{3, 0},
			[]int{4, 0},
			[]int{5, 0},
		},
		expected: []int{
			0,
			2,
			2,
			3,
			4,
			5,
		},
	},
	{
		skipped: false,
		input: [][]int{
			[]int{2, 8},
			[]int{0, 0},
			[]int{0, 0},
			[]int{0, 0},
			[]int{1, 0},
			[]int{1, 0},
			[]int{1, 1},
			[]int{1, 2},
			[]int{1, 3},
		},
		expected: []int{
			0,
			0,
			0,
			1,
			1,
			1,
			2,
			-1,
		},
	},
	{
		skipped: false,
		input: [][]int{
			[]int{2, 8},
			[]int{0, 0},
			[]int{0, 0},
			[]int{0, 0},
			[]int{1, 1},
			[]int{1, 0},
			[]int{1, 0},
			[]int{1, 2},
			[]int{1, 3},
		},
		expected: []int{
			0,
			0,
			0,
			1,
			2,
			-1,
			-1,
			-1,
		},
	},
	{
		skipped: false,
		input: [][]int{
			[]int{1, 5},
			[]int{999999, 1},
			[]int{1000000, 0},
			[]int{1000000, 1},
			[]int{1000000, 0},
			[]int{1000000, 0},
		},
		expected: []int{
			999999,
			1000000,
			1000000,
			-1,
			-1,
		},
	},
}
