package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func characterFreqs(words string) {
	wordByWord := strings.Split(words, " ")

	var charCountMap = make(map[string]int)
	var wordChan = make(chan string, len(charCountMap))

	for _, word := range wordByWord {		
		go countChar(&charCountMap, wordChan)

		wordChan <- word
	}

	for keyChar, charCount := range charCountMap {
		fmt.Println(keyChar, ":", charCount)
	}
}

func countChar(charCountMap *map[string]int, word chan string)  {

	for _, char := range <-word {
		if _, found := (*charCountMap)[string(char)]; found {
			(*charCountMap)[string(char)]++
			continue
		}

		(*charCountMap)[string(char)] = 1
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	words, _ := reader.ReadString('\n')
	words = strings.ReplaceAll(words, "\n", "")
	characterFreqs(words)
}