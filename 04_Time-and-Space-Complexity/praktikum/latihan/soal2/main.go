package main

import "fmt"

func Power(x int, y int) int {
	if y == 0 {
		return 1
	}

	if y % 2 == 0 {
		return Power(x, y / 2) * Power(x, y / 2)
	}

	return x * Power(x, y / 2) * Power(x, y / 2)
}

func main() {
	fmt.Println(Power(2, 6))
}

