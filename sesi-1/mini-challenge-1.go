package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func miniChallengeOne() {
	fmt.Print("Input angka lebih besar dari 1: ")

	reader := bufio.NewReader(os.Stdin)

	myNumber, _ := reader.ReadString('\n')

	myNumber = strings.Replace(myNumber, "\n", "", -1)

	convertedNumber, _ := strconv.Atoi(myNumber)

	// Setiap looping dengan angka kelipatan 3 maka diganti dengan kata “Fizz”
	// Setiap looping dengan angka kelipatan 5 maka diganti dengan kata “Buzz”
	// Setiap looping dengan angka kelipatan 3 dan 5 maka diganti dengan kata “FizzBuzz”

	for i := 1; i <= convertedNumber; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)

		}
	}

}
