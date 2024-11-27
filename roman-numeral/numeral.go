package roman_numeral

import (
	"errors"
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

var ErrValueTooLarge = errors.New("input number can't exceed 3999")
var ErrNoRomanZero = errors.New("invalid Roman numeral: no equivalent Arabic value")

func ConvertToRoman(arabic uint16) (string, error) {
	if arabic > 3999 {
		return "", ErrValueTooLarge
	}

	var res strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			res.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return res.String(), nil
}

func ConvertToArabic(roman string) (uint16, error) {
	var arabic uint16
	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	if arabic <= 0 {
		return 0, ErrNoRomanZero
	}
	return arabic, nil
}
