package main

import (
	"fmt"
	"sort"
)


func findElement(elements []int, element int) int {
	sort.Slice(elements, func(i, j int) bool {
		return elements[i] < elements[j]
	})

	left := 0
	right := len(elements) - 1
	index := -1

	for left <= right {
		middle := (left + right) / 2
		if element < elements[middle] {
			right = middle - 1
		} else if element > elements[middle] {
			left = middle + 1
		} else {
			index = middle
			break
		}
	}
	
	return index
}

func main() {
	var elements = []int{10, 3, 31, 23, 100, 11, 1}

	fmt.Println(findElement(elements, 1000))
}