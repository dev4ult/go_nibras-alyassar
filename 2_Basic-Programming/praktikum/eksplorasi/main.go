package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Apakah Palindrome?")

	var word string

	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Masukan kata: ")
    word, _ = scanner.ReadString('\n')
	word = strings.ReplaceAll(word, "\r\n", "")
	fmt.Println("Captured:", word)
	ValidPalindrome(word)
}

func ValidPalindrome(word string) {
	var backWord string
	
	for i := len(word) - 1; i >= 0; i-- {
		backWord += string(word[i])
	}

	if strings.TrimLeft(word, " ") == backWord {
		fmt.Println("a Palindrome")
	} else {
		fmt.Println("Not a Palindrome")
	}
}