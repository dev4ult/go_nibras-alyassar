package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func SelisihDiagonalAbs() float64 {
	scanner := bufio.NewReader(os.Stdin)

	rowOne, _ := scanner.ReadString('\n') 
	rowTwo, _ := scanner.ReadString('\n') 
	rowThree, _ := scanner.ReadString('\n')
	
	
	
	var diagonalArr [][]string

	diagonalArr = append(diagonalArr, strings.Split(rowOne, " "))
	diagonalArr = append(diagonalArr, strings.Split(rowTwo, " "))
	diagonalArr = append(diagonalArr, strings.Split(rowThree, " "))

	var result int = 0

	fmt.Println(diagonalArr[0][2])
	
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			if j == 2 {
				diagonalArr[i][j] = strings.ReplaceAll(diagonalArr[i][j], "\n", "")
			}

			value, _ := strconv.Atoi(diagonalArr[i][j])

			if i == j {
				result += value
			}

			if i + j == 2 {
				result -= value
			}
		}
	}

	return math.Abs(float64(result))
	
}

func main() {
	SelisihDiagonalAbs()
}