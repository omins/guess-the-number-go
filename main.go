package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/fatih/color"
)


func main() {
	printGreeting()
	scanner := bufio.NewScanner(os.Stdin)
	
	color.Blue("\nBefore starting, please enter maximum number of the range:")
	maxNumber := getNumberInput(scanner, -1)

	fmt.Printf("\nGreat! let's start the game :)\nI have a number in my mind from 0 to %d\n\n", maxNumber)
	answer := getRandomNumber(0, maxNumber)
	tryCount := playGame(scanner, answer)
	
	printResult(tryCount)
}

func printGreeting() {
	color.Green("---------------------------------------\n")
	color.Green("| Welcome to the Number Guessing Game |\n")
	color.Green("---------------------------------------\n")
}

func getNumberInput(scanner *bufio.Scanner, base int) int {
	for {
		scanner.Scan()
		input := scanner.Text()
		maxNumber, err := strconv.Atoi(input)
		
		if err == nil && maxNumber > base {
			return maxNumber
		}

		color.Red("Please enter a valid number greater than %d", base)
	}
}

func getRandomNumber(min, max int) int {
	return rand.Intn(max - min + 1) + min
}

func playGame(scanner *bufio.Scanner, answer int) (tryCount int) {
	for {
		tryCount++
		color.Blue("Please enter your guess:")
		guess := getNumberInput(scanner, -1)

		if guess == answer {
			break
		} else if guess > answer {
			color.Red("The number is smaller than %d\n", guess)
		} else {
			color.Red("The number is greater than %d\n", guess)
		}
	}
	return
}

func printResult(tryCount int) {
	switch tryCount {
		case 1: 
			color.Yellow("\nWow! You found the number in your first try!")
		default:
			color.Yellow("\nCongratulations! You found the number in %d tries", tryCount)
	}
}