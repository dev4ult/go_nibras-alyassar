package main

import "fmt"

func RomanNumerals(number int) string {

	var romanNumber string = ""


	romanTen := number / 10 
	if romanTen != 0 {
		for i := 0; i < romanTen; i++ {
			romanNumber += "X"
		}
	}

	romanFive := number % 10 / 5
	if romanFive != 0 {
		for i := 0; i < romanFive; i++ {
			romanNumber += "V"
		}
	}

	return romanNumber
}

func main() {

	fmt.Println(RomanNumerals(25))
}