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

	// fmt.Println(caesar(3, "abc")) // def

	// fmt.Println(caesar(2, "alta")) // def

	// fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

	// fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

	resultASCII := 1 + 1000

	for resultASCII > 26 {
		fmt.Print(resultASCII, " ")
		resultASCII = 1 + (resultASCII % 26)
	}

	

	// for resultASCII > 26 {
	// 	fmt.Print(resultASCII, " ")
	// 	resultASCII = 96 + (resultASCII - 26)
	// }

	// fmt.Println(resultASCII)

	// fmt.Println("============")


	// resultASCII = 2 + 27
	// for resultASCII > 26 {
	// 	fmt.Print(resultASCII, " ")
	// 	resultASCII = resultASCII - 26
	// }

	// fmt.Println(resultASCII)


}
