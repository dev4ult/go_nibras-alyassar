package main

import "fmt"

func main() {
	TampilkanSegitigaAsterisk(5)
}

func TampilkanSegitigaAsterisk(jumlahSisi int) {
	for i := 0; i < jumlahSisi; i++ {
		for j := 0; j < jumlahSisi - i - 1; j++ {
			fmt.Print(" ")
		}

		for k := 0; k <  2 * (i + 1) - 1; k++ {
			if k % 2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		} 
		fmt.Println(" ")
	}
}