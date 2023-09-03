package main

import (
	"fmt"
	"strconv"
)

func ConvDecToBinary(n int) string {

	var binaryElements string

	for i := n; i > 0; i /= 2 {
		binaryElements += strconv.Itoa(i % 2)
	}

	var reversedStr string = ""

	for i := len(binaryElements) - 1; i >= 0; i-- {
		reversedStr += string(binaryElements[i])
	}

	return reversedStr
}

func main() {
	fmt.Println(ConvDecToBinary(12))
}