package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func miniChallengeTwo() {

	fmt.Print("Input kata/kalimat: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	input = strings.Replace(input, "\n", "", -1)

	countPerChar := make(map[string]int)

	charSlice := stringToCharSlice(input)

	// fmt.Println("Hasil dari: ", charSlice)

	for _, value := range charSlice {
		fmt.Println(value)
		countPerChar[value]++
	}

	fmt.Println("Count per char: ", countPerChar)
}

func stringToCharSlice(s string) []string {
	runes := []rune(s)
	result := make([]string, len(runes))
	for i, r := range runes {
		result[i] = string(r)
	}
	return result
}
