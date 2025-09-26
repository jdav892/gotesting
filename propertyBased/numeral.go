package main

import (
	"strings"
)

type RomanNumeral struct {
	Value int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{50, "L"},
	{49, "XLIX"},
	{47, "XLVII"},
	{40, "XL"},
	{39, "XXXIX"},
	{20, "XX"},
	{18, "XVIII"},
	{14, "XIV"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}
