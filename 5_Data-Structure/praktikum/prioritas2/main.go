package main

import "fmt"

func PairSum(arr []int, target int) []int {
	var indexArr []int

	for index, val := range arr {
		if val > target {
			continue
		}

		for index2, val2 := range arr {
			if index2 == index {
				continue
			}
			
			total := val + val2 

			if total == target {
				indexArr = append(indexArr, index)
				indexArr = append(indexArr, index2)
				break
			}
		} 

		if len(indexArr) > 0 {
			break
		}
	}

	return indexArr
}


func main() {

	// Test cases

	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]

	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11)) // [0, 2]

	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // [2, 3]

	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10)) // [1, 2]

	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6)) // [0, 1]

}

