package main

import "fmt"

func main() {
	GanjilAtauGenap(10)
}

func GanjilAtauGenap(sampaiAngka int) {
	for i := 1; i <= sampaiAngka; i++ {
		if i%2 == 0 {
			fmt.Println(i, " adalah bilangan Genap")
		} else {
			fmt.Println(i, " adalah bilangan Ganjil")
		}
	}
}