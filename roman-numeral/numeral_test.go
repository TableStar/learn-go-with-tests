package roman_numeral

import (
	"errors"
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic             uint16
	Roman              string
	ErrorRomanToArabic error
	ErrorArabicToRoman error
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
	{Arabic: 5000, Roman: "", ErrorRomanToArabic: ErrValueTooLarge, ErrorArabicToRoman: ErrNoRomanZero},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got, err := ConvertToRoman(test.Arabic)
			if test.ErrorRomanToArabic != nil {
				if err == nil {
					t.Fatalf("expected error but got nil for Arabic %d", test.Arabic)
				}
				if !errors.Is(err, test.ErrorRomanToArabic) {
					t.Fatalf("unexpected error: got %v, want %v", err, test.ErrorRomanToArabic)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error for Arabic %d: %v", test.Arabic, err)
			}
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertToArabs(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got, err := ConvertToArabic(test.Roman)
			if test.ErrorArabicToRoman != nil {
				if err == nil {
					t.Fatalf("expected error but got nil for Roman %q", test.Roman)
				}
				if !errors.Is(err, test.ErrorArabicToRoman) {
					t.Fatalf("unexpected error: got %v, want %v", err, test.ErrorArabicToRoman)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error for Roman %q: %v", test.Roman, err)
			}
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			log.Println(arabic)
			return true
		}
		roman, _ := ConvertToRoman(arabic)
		fromRoman, _ := ConvertToArabic(roman)
		return fromRoman == arabic
	}
	if err := quick.Check(assertion, &quick.Config{MaxCount: 100}); err != nil {
		t.Error("failed checks", err)
	}
}
