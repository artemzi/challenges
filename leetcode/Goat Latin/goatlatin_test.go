package goatlatin

import (
	"testing"
)

func TestToGoatLatin(t *testing.T) {
	input := "I speak Goat Latin"
	expected := "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"

	if r := toGoatLatin(input); r != expected {
		t.Errorf("\nResult for \n{%s}\n is wrong. got: \n{%s}\n, expected: \n{%s}\n", input, r, expected)
	}

	input = "The quick brown fox jumped over the lazy dog"
	expected = "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"

	if r := toGoatLatin(input); r != expected {
		t.Errorf("\nResult for \n{%s}\n is wrong. got: \n{%s}\n, expected: \n{%s}\n", input, r, expected)
	}
}
