package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

var length = 3
var guesses = 10

func main() {
	displayHelp()
	start()
}

func displayHelp() {
	fmt.Println("Bagels deduction game")
	fmt.Printf("I'm thinking of a %d-digit number. Try to guess what it is\n", length)
	fmt.Println("The clues i give are...")
	fmt.Println("When I say:    That means:")
	fmt.Println("  Bagels       None of the digits is correct.")
	fmt.Println("  Pico         One digit is correct but in the wrong position.")
	fmt.Println("  Fermi        One digit is correct and in the right position.")
}

func start() {
	fmt.Printf("I have thought up a number. You have %d guesses to get it.\n", guesses)

	number := randomNumber()

	for guesses > 0 {
		guess := takeGuess()
		guesses--

		if number == guess {
			fmt.Print("You got it!")
			break
		} else if guesses == 0 {
			fmt.Printf("You have lost! The number was %d\n", number)
		} else {
			numberDigits := toDigits(number)
			guessDigits := toDigits(guess)

			clues := []string {}

			for i, digit := range guessDigits {
				if digit == numberDigits[i] {
					clues = append(clues, "Fermi")
				} else {
					for _, other := range numberDigits {
						if digit == other {
							clues = append(clues, "Pico")
						}
					}
				}
			}

			if len(clues) == 0 {
				clues = append(clues, "Bagels")
			}

			sort.Strings(clues)
			fmt.Println(strings.Join(clues, " "))
		}
	}
}

func randomNumber() int {
	rand.Seed(time.Now().UnixNano())
	digits := rand.Perm(10)
	var number int
	for i := 0; i < length; i++ {
		number *= 10
		number += digits[i]
	}
	return number
}

func toDigits(number int) []int {
	digits := make([]int, length)
	for i, char := range strconv.Itoa(number) {
		digits[i] = int(char - '0')
	}
	return digits
}

func takeGuess() int {
	var guess string
	fmt.Scanln(&guess)

	for len(guess) != length {
		fmt.Printf("You have to provide %d-digit number\n", length)
		fmt.Scanln(&guess)
	}

	result, err := strconv.Atoi(guess)

	for err != nil {
		fmt.Printf("You have to provide %d-digit number\n", length)
		fmt.Scanln(&guess)
		result, err = strconv.Atoi(guess)
	}

	return result
}