package main

import "fmt"

func main() {
	PrintFaktor(26)
}

func PrintFaktor(angka int) {
	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			fmt.Println(i)
		}
	}
}