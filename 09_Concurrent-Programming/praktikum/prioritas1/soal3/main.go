package main

import "fmt"

func MultipleWithBuffChannel(x chan int) int {
	valX := <-x
	for i := 1; i > 0; i++ {
		if valX*1 < 0 {
			i = 0
		}
		fmt.Println(valX * i)
	}

	return -1
}

func main() {
	number := make(chan int, 1)
	number <- 3
	MultipleWithBuffChannel(number)
}
