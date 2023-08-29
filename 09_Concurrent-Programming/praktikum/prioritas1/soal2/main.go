package main

import "fmt"

func MultipleWithChannel(x chan int) int {
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
	number := make(chan int)
	go MultipleWithChannel(number)
	number <- 3
}