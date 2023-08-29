package main

import (
	"fmt"
	"time"
)

func main() {
	theMine := [5]string{"ore1", "ore2", "ore3"}
	oreChan := make(chan string)

	go func(mine [5]string) {
		for _, ore := range mine {
			oreChan <- ore
		}
	}(theMine)

	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChan
			fmt.Println("Miner : Received " + foundOre + " from finder")
		}
	}()
	<-time.After(time.Second * 5)
}
