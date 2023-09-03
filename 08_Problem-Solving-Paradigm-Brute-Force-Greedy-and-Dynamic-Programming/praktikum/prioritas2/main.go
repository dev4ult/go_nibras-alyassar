package main

import (
	"fmt"
	"math"
)


func NonGreedyFrog(index int, jumps []int) float64 {
	var prevJump, prevJump2, minimumTotalStep int = 0, 0, 0
	var oneStep, twoStep float64 = 0, 0

	for i := 1; i < len(jumps); i++ {
		oneStep = float64(prevJump) + math.Abs(float64(jumps[i] - jumps[i - 1]))

		if i > 1 {
			twoStep = float64(prevJump2) + math.Abs(float64(jumps[i] - jumps[i - 2]))
		}

		if twoStep < oneStep {
			minimumTotalStep = int(twoStep)
		} else {
			minimumTotalStep = int(oneStep)
		}

		prevJump2 = prevJump
		prevJump = minimumTotalStep
	}

	return float64(prevJump)
	// untuk initialisasi pertama agar tidak error ketika index <= 1

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

	fmt.Println(GreedyFrog([]int{10, 30, 40, 20})) // 30

	var steps = []int{30, 10, 60, 10, 60, 50} 
	fmt.Println(NonGreedyFrog(len(steps) - 1, steps)) // 40

}