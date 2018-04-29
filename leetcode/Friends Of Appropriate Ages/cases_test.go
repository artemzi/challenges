package friendsages

type goatlatinTest struct {
	input    []int
	expected int
}

var stringTestCases = []goatlatinTest{
	{
		input:    []int{16, 16},
		expected: 2,
	},
	{
		input:    []int{16, 17, 18},
		expected: 2,
	},
	{
		input:    []int{20, 30, 100, 110, 120},
		expected: 0,
	},
}
