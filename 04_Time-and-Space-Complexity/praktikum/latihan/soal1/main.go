package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(primeNumber(1000000007))
}

func primeNumber(number int) string {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number % i == 0 {
			return "Bilangan Prima"
		}
	}

	return "Bukan Bilangan Prima"
}
