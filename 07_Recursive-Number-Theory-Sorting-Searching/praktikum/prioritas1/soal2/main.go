package main

import (
	"fmt"
	"sort"
)


type pair struct {
	name string
	count int
}


func MostAppearItem(items []string) []pair {
	var list []pair

	for _, itemName := range items {

		isExist := false
		
		for index, aListItem := range list {
			if itemName == aListItem.name {
				list[index].count++
				isExist = true
				break
			}			
		}

		if isExist {
			continue
		}

		list = append(list, pair{
			name: itemName,
			count: 1,
		})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].count < list[j].count
	})
	
	return list
}


func main() {

	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) // golang->1 ruby->2 js->4

	for _, list := range pairs {

		fmt.Print(list.name, " -> ", list.count, " ")

	}

	fmt.Println()


	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}) // C->1 D->2 B->3 A->4

	for _, list := range pairs {

		fmt.Print(list.name, " -> ", list.count, " ")

	}

	fmt.Println()


	pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) // football->1 basketball->1 tenis->1

	for _, list := range pairs {

		fmt.Print(list.name, " -> ", list.count, " ")

	}

}

