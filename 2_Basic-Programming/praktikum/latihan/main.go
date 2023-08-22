package main

import (
	"fmt"
)

func main() {
	VariableExample()
	
	OperatorNExpression()

	StringConcatenation()

	Branching()

	Looping()
}

func VariableExample() {
	var isLogin bool = false
	fmt.Println(isLogin)

	var prime int = 7
	fmt.Println(prime)

	var decimal float64 = 3.5
	fmt.Println(decimal)

	const pi float64 = 3.14
	fmt.Println(pi)

	// short variable definition
	name := "Nibras"
	fmt.Println(name)
}

func OperatorNExpression() {
	x := 1 + 2
	fmt.Println(x)
}

func TriangleArea() {
	a := 12
	t := 10
	L := (a * t) / 2
	fmt.Println("Luas Segitiga : ", L)
}

func StringConcatenation() {
	str := "Nibras" + " " + "Alyassar"
	fmt.Println(str)
}

func Branching() {
	hour := 15

	// basic branching with conditional if else
	if (hour < 12) {
		fmt.Println("Good Morning")
	} else if (hour >= 12 && hour <= 18) {
		fmt.Println("Good Evening")
	} else {
		fmt.Println("Good Afternoon")
	}

	// switch case does not need a break in golang
	today := 2
	switch(today) {
		case 1:
			fmt.Println("It's Sunday")
		case 2:
			fmt.Print("It's Monday")
		case 3:
			fmt.Println("It's Tuesday")
		case 4:
			fmt.Print("It's Wednesday")
		default :
			fmt.Println("Unknown day")
	}
}

func Looping() {

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if (i == 3) {
			continue
		} else if (i == 4) {
			break
		}
	}

	// looping over string
	str := "Nibras"
	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[i]), "-")
	}

	for pos, char := range str {
		fmt.Printf("Character %c start from position %d\n", char, pos)
	}
}