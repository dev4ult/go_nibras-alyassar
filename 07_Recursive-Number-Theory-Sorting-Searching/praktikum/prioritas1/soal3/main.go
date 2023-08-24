package main

import "fmt"

func primeX(number int) int {
	var primeNumber int = 1
	counter := 1

	for counter <= number {
		primeNumber++

		isPrime := true
		for i := 2; i < primeNumber; i++ {
			if primeNumber % i == 0 {
				isPrime = false
			}
		}

		if isPrime {
			counter++
		}
	}

	return primeNumber
}

func main() {

	fmt.Println(primeX(1)) // 2

	fmt.Println(primeX(5)) // 11

	fmt.Println(primeX(8)) // 19

	fmt.Println(primeX(9)) // 23

	fmt.Println(primeX(10)) // 29

}