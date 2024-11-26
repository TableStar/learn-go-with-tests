package roman_numeral

import "strings"

func ConvertToRoman(arabic int) string {
	var res strings.Builder

	for arabic > 0 {
		switch {
		case arabic > 4:
			res.WriteString("V")
			arabic -= 5
		case arabic > 3:
			res.WriteString("IV")
			arabic -= 4
		default:
			res.WriteString("I")
			arabic--
		}
	}
	return res.String()
}
