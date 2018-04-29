package goatlatin

type goatlatinTest struct {
	input    string
	expected string
}

var stringTestCases = []goatlatinTest{
	{
		input:    "I speak Goat Latin",
		expected: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
	},
	{
		input:    "The quick brown fox jumped over the lazy dog",
		expected: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
	},
}
