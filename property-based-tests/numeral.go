package numeral

import (
	"strings"
)

var romanLetters = []string{
	"M", "CM", "D", "CD",
	"C", "XC", "L", "XL",
	"X", "IX", "V", "IV",
	"I",
}

var romanLetterValues = []int{
	1000, 900, 500, 400,
	100, 90, 50, 40,
	10, 9, 5, 4,
	1,
}

func ConvertToRoman(num int) string {
	var str strings.Builder

	for i, value := range romanLetterValues {
		for num >= value {
			str.WriteString(romanLetters[i])
			num -= value
		}

		if num == 0 {
			break
		}
	}

	return str.String()
}

func ConvertToNumber(roman string) int {
	number := 0

	for i, letters := range romanLetters {
		for strings.HasPrefix(roman, letters) {
			number += romanLetterValues[i]
			roman = strings.TrimPrefix(roman, letters)
		}

		if roman == "" {
			break
		}
	}

	return number
}
