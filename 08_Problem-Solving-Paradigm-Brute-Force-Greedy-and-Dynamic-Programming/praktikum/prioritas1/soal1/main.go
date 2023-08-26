package main

import (
	"fmt"
	"strconv"
)

func loopBinary(n int) []string {

	var binaryElements []string

	for i := 0; i <= n; i++ {
		binaryElements = append(binaryElements, strconv.FormatInt(int64(i), 2))
	}

	return binaryElements
}

func main() {
	fmt.Println(loopBinary(4))
}