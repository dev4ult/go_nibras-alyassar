package main

import (
	"fmt"
	"math"
	"sort"
)


func Frog(jumps []int) int {

	var totalOptimum []int

	sort.Ints(totalOptimum)
	return totalOptimum[len(totalOptimum) - 1]

}


// Nyobain metode greedy (tidak selalu optimum)
func GreedyFrog(jumps []int) int {
	var totalOptimum int = 0

	for i := 0; i < len(jumps) - 1; i++ {		
		oneStep := math.Abs(float64(jumps[i + 1]) - float64(jumps[i]))
		twoSteps := math.Abs(float64(jumps[i + 2]) - float64(jumps[i]))
		
		if twoSteps > oneStep {
			totalOptimum += int(twoSteps)
			i += 2
			continue
		}

		if oneStep > twoSteps {
			totalOptimum += int(oneStep)
			i++
			continue
		}
	}

	return totalOptimum
}


func main() {

	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30

	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40

}