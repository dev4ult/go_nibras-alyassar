package main

import (
	"fmt"
	"sync"
)

func PrintIndex() int {
	var i int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		i = 5
		
		wg.Done()
	}()
	wg.Wait()

	return i
}

func main() {
	fmt.Println(PrintIndex())
}