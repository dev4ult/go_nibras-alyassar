package main

import (
	"fmt"
	"math"
)

func PrimeNumber(number int) string {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number % i == 0 {
			return "Bukan Bilangan Prima"
		}
	}

	return "Bilangan Prima"
}

func main() {
	fmt.Println(PrimeNumber(503))
}


