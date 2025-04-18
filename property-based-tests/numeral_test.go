package numeral

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	Number int
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestConvertToRoman(t *testing.T) {
	for _, test := range cases {
		testname := fmt.Sprintf("converts %v to %v", test.Number, test.Roman)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, test.Roman, ConvertToRoman(test.Number))
		})
	}
}

func TestConvertToNumber(t *testing.T) {
	for _, test := range cases {
		testname := fmt.Sprintf("converts %v to %v", test.Roman, test.Number)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, test.Number, ConvertToNumber(test.Roman))
		})
	}
}
