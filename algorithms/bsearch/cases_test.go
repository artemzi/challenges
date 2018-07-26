package bsearch

type (
	inputData struct {
		data []int
		val  int
	}

	bsearchTest struct {
		input    *inputData
		expected int
	}
)

var stringTestCases = []bsearchTest{
	{
		input: &inputData{
			[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
			101,
		},
		expected: -1,
	},
	{
		input: &inputData{
			[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
			5,
		},
		expected: 2,
	},
	{
		input: &inputData{
			[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
			41,
		},
		expected: 12,
	},
	{
		input: &inputData{
			[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
			83,
		},
		expected: 22,
	},
}
