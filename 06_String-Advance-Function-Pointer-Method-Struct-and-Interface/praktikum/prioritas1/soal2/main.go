package main

import "fmt"

// Properti name dan score dalam bentuk slice?
type Student struct {
	name  []string
	score []float32
}

func (s Student) PrintFinalEvaluation() {
	if len(s.name) != len(s.score) {
		fmt.Println("Jumlah nama dan score siswa tidak sesuai")
		return
	}

	var average float32
	var maxScoreIndex, minScoreIndex int = 0, 0
	var maxScore, minScore float32 = s.score[0], s.score[0]
	
	for index, currentScore := range s.score {
		average += currentScore

		if currentScore > maxScore {
			maxScoreIndex = index
		}

		if currentScore < minScore {
			minScoreIndex = index
		}
	}

	average /= float32(len(s.score))

	fmt.Println("Average Score :", average)
	fmt.Println("Min Score of Students :", s.name[minScoreIndex], fmt.Sprintf("(%d)", int(minScore)))
	fmt.Println("Max Score of Students :", s.name[maxScoreIndex], fmt.Sprintf("(%d)", int(maxScore)))
}

func main() {
	classB := Student {
		name: []string{"Rizky", "Kobar", "Ismail", "Umam", "Yopan"},
		score: []float32{80, 75, 70, 75, 60},
	}

	classB.PrintFinalEvaluation()
}