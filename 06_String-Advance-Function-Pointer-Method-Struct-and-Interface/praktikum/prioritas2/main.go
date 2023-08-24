package main

import "fmt"

func caesar(offset int, input string) string {

	var result string = ""

	for _, charASCII := range input {
		resultASCII := charASCII + int32(offset) 
		for resultASCII > 122 {
			resultASCII = resultASCII - 26
		}

		result += string(resultASCII)
	}
	
	return result

}

func main() {

	fmt.Println(caesar(3, "abc")) // def

	fmt.Println(caesar(2, "alta")) // def

	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl


}
