package main

import "fmt"

func PascalTriangle(numRows int) [][]int {

	var pascal [][]int

	for i := 1; i <= numRows; i++ {
		pascal = append(pascal, []int{})
		for j := 1; j <= i; j++ {
			if j == 1 || j == i {
				pascal[i-1] = append(pascal[i-1], 1)
				continue
			}
			pascal[i-1] = append(pascal[i-1], pascal[i-2][j-1] + pascal[i-2][j-2])
		}
	}

	return pascal
}

func main() {

	fmt.Println(PascalTriangle(5))

}
