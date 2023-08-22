package main

import "fmt"

func main() {
	fmt.Println(primeNumber(17))
}


// Time Complexity = O(n)
func primeNumber(n int) bool {
	var ctr int = 0
	for i := 1; i <= n; i++ {
		if n % i == 0 {
			ctr++
		}
	}

	return ctr == 2
}

// Time Complexity = O(n / 2)
func primeNumber2(n int) bool {
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

func primeNumber3(n int) bool {
	
}
