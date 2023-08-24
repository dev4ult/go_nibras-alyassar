package main

import (
	"fmt"
	"strings"
)


func Compare(a, b string) string {
	if len(a) < len(b) && strings.Contains(b, a) {
		return a
	}

	if len(a) > len(b) && strings.Contains(a, b) {
		return b
	}

	return "There is no word match between " + a + " and " + b
}


func main() {

	fmt.Println(Compare("AKA", "AKASHI")) //AKA

	fmt.Println(Compare("KANGOORO", "KANG")) //KANG

	fmt.Println(Compare("KI", "KIJANG")) //KI

	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU

	fmt.Println(Compare("ILALANG", "ILA")) //ILA

}