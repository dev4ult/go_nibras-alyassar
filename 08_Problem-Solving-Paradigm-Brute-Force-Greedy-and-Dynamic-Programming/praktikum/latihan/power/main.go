package main

import "fmt"

func power(x int, y int) int {
	if y == 0 {
		return 1
	}

	if y % 2 == 0 {
		return power(x, y / 2) * power(x, y / 2)	
	}
	
	return x * power(x, y / 2) * power(x, y / 2)
}

func main() {

	fmt.Println(power(2, 8))

}