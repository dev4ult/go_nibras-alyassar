package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	
	var arrMap = make(map[int]int)
	
	for _, perAngka := range angka {
		key, _ := strconv.Atoi(string(perAngka))
		_, exist := arrMap[key]

		if exist {
			arrMap[key]++
		} else {
			arrMap[key] = 1
		}
	}

	var result []int

	for key, value := range arrMap {
		if value == 1 {
			result = append(result, key)
		}
	}


	return result

}

func main() {

	// Test cases

	fmt.Println(munculSekali("1234123")) // [4]

	fmt.Println(munculSekali("76523752")) // [6 3]

	fmt.Println(munculSekali("12345")) // [1 2 3 4 5]

	fmt.Println(munculSekali("1122334455")) // []

	fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]

}