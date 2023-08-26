package main

import "fmt"

type conversion struct {
	value int
	digit string
}

func RomanNumerals(number int) string {
		
	conversions := []conversion{
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

	var romanNumber string = ""

	for _, obj := range conversions {
		for obj.value <= number {
			number -= obj.value
			romanNumber += obj.digit
		}
	}
	
	return romanNumber
}

func main() {

	fmt.Println(RomanNumerals(1646))
}