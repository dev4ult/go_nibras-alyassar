package main

import (
	"fmt"
	"sync"
)

type SafeNumber struct {
	val int
	m sync.Mutex
}

func (safeNumber *SafeNumber) Set(val int) {
	safeNumber.m.Lock()
	defer safeNumber.m.Unlock()

	safeNumber.val = val
}

func (safeNumber *SafeNumber) Get() int {
	safeNumber.m.Lock()
	defer safeNumber.m.Unlock()

	return safeNumber.val
}

func main() {
	i := &SafeNumber{}
	go func() {
		i.Set(5)
	}()
	// time.Sleep(1 * time.Second)
	fmt.Println(i.Get())
}