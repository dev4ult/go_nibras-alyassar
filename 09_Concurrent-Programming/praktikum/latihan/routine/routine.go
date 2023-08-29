package main

import (
	"fmt"
	"time"
)

func hello(s string) {
	for _, char := range s {
		fmt.Print(string(char))
	}
}

func main() {
	go hello("nibras alyassar")
	time.Sleep(1 * time.Second)
}