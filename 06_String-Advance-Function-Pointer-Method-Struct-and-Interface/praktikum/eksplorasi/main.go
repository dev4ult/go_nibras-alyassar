package main

import "fmt"

type student struct {
	name string
	nameEncode string
	score int
}

type Chiper interface {
	Encode() string
	Decode() string
}


func GetSortedIndex(chr string) int {
	if len(chr) > 1 {
		return -1
	}

	lowercaseAlphabet := "abcdefghijklmnopqrstuvwxyz"
	uppercaseAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// indexOutput := 0
	for index, lwrChar := range lowercaseAlphabet {
		if chr == string(lwrChar) || chr == string(uppercaseAlphabet[index]) {
			return index
		}
	}

	// if not found
	return -1
}

func GetReversedIndex(chr string) int {
	if len(chr) > 1 {
		return -1
	}

	lowercaseAlphabet := "zyxwvutsrqponmlkjihgfedcba"
	uppercaseAlphabet := "ZYXWVUTSRQPONMLKJIHGFEDCBA"

	// indexOutput := 0
	for index, lwrChar := range lowercaseAlphabet {
		if chr == string(lwrChar) || chr == string(uppercaseAlphabet[index]) {
			return index
		}
	}

	// if not found
	return -1
}

func (s *student) Encode() string {
	reversedAlphabet := "zyxwvutsrqponmlkjihgfedcba"
	
	var nameEncode = ""

	for _, char := range s.name {
		indexChar := GetSortedIndex(string(char))

		if indexChar != -1 {
			nameEncode += string(reversedAlphabet[indexChar])
		}
	}

	return nameEncode
}


func (s *student) Decode() string {
	sortedAlphabet := "abcdefghijklmnopqrstuvwxyz"

	var nameDecode = ""

	for _, char := range s.name {
		indexChar := GetReversedIndex(string(char))

		if indexChar != -1 {
			nameDecode += string(sortedAlphabet[indexChar])
		}
	}


	return nameDecode
}


func main() {

	var menu int

	var a student = student{}

	var c Chiper = &a


	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")

	fmt.Scan(&menu)


	if menu == 1 {

		fmt.Print("\nInput Student Name: ")

		fmt.Scan(&a.name)

		fmt.Print("\nEncode of student's name " + a.name + "is : " + c.Encode())

	} else if menu == 2 {

		fmt.Print("\nInput Student Name: ")

		fmt.Scan(&a.name)

		fmt.Print("\nDecode of student's name " + a.name + "is : " + c.Decode())

	}

}