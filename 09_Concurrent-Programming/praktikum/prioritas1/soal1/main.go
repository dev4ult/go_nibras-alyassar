package main

import (
	"fmt"
	"time"
)

func Multiple(x int) int {
	for i := 1; i > 0; i++ {
		if x*1 < 0 {
			return 1
		}
		fmt.Println(x * i)
	}

	return -1
}

func main() {
	go Multiple(12)
	time.Sleep(3 * time.Second)
}

