package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	
	for _, itemB := range arrayB {
		continueLoop := false
		for _, itemA := range arrayA {
			if itemA == itemB {
				continueLoop = true		
			}
		}

		if continueLoop {
			continue
		}
		
		arrayA = append(arrayA, itemB)
	}

	return arrayA
}

func main() {
	// Test cases

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))

	// []

}