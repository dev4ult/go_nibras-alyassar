package main

import "fmt"

func main() {
	fmt.Println(pow(2, 5))
}

func pow(x int, y int) int {
	if y == 0 {
		return 1
	}

	if y % 2 == 0 {
		return x * pow(x, y / 2)
	} else {
		return x * pow(x, y / 2) * pow(x, y / 2)

	}
}