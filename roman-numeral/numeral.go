package roman_numeral

import "strings"

func ConvertToRoman(arabic int) string {
	var res strings.Builder

	for i := arabic; i > 0; i-- {
		if i == 5 {
			res.WriteString("V")
			break
		}
		if i == 4 {
			res.WriteString("IV")
			break
		}
		res.WriteString("I")
	}
	return res.String()
}
