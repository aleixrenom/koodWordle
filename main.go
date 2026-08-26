package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"unicode"
	// "encoding/csv"
	// "io"
)

// --------------------------------------------------------------
// PACKAGE VARIABLES
// --------------------------------------------------------------
	
const (
	Reset = "\033[0m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	White = "\033[37m"
)

type GameData struct {
	username	string
	secretWord	string
	attempts	int
	victory		bool
}

var remainingLetters []rune = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
var wordToGuess []rune
var currentStats GameData = GameData{}


func main() {
	// --------------------------------------------------------------
	// ARGUMENT VALIDATION
	// --------------------------------------------------------------

	args := os.Args[1:] // grab all user-given arguments (first one is always just path)

	if len(args) != 1 {
		fmt.Println("Please provide a number as command line argument")
		os.Exit(0)
	}

	wordNumber, wErr := strconv.Atoi(args[0])
	if wErr != nil {
		fmt.Println("Invalid command-line argument. Please launch with a valid number.")
		os.Exit(0)
	}

	if wordNumber <= 0 {
		// fmt.Println("Please, give a positive number as the first argument.")
		fmt.Println("Enter your username: Invalid word number.")
		fmt.Println("Press Enter to exit...")
		os.Exit(0)
	}

	wordToGuess = getWordToGuess(wordNumber)
	//fmt.Printf("(Don't tell anyone, but the word is \"%s\")\n", string(wordToGuess))

	// --------------------------------------------------------------
	// GAME
	// --------------------------------------------------------------

	attemptsRemaining := 6

	scanner := bufio.NewScanner(os.Stdin)

	// Ask for username and store it
	for currentStats.username == "" {
		fmt.Print("Enter your username: ")
		if scanner.Scan() {
			currentStats.username = strings.TrimSpace(scanner.Text())
		}
	}

	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")
	
	// Main game loop
	for {
		guess := ""
		fmt.Print("Enter your guess:")
		if scanner.Scan() {
			guess = strings.TrimSpace(scanner.Text())
		} else {
			// fmt.Printf("There was a problem gathering input.")
			os.Exit(0)
		}

		// Input validation
		if !wordExists(guess) {
			fmt.Println("Please enter a valid word.")
			continue
		}
		if len(guess) != 5 {
			fmt.Println("Your guess must be exactly 5 letters long.")
			continue
		}
		con := false
		for _, ch := range guess {
			if unicode.IsDigit(ch) || unicode.IsUpper(ch) {
				fmt.Println("Your guess must only contain lowercase letters.")
				con = true
				os.Exit(0)
			}
		}
		if con { continue }

		// Correct guess
		if guess == string(wordToGuess) {
			fmt.Println("Congratulations! You've guessed the word correctly.")
			currentStats.victory = true
			break
		}

		// If wrong guess, print the feedback
		printFeedback(guess)

		// Print remaining letters
		fmt.Print("\nRemaining letters: ")
		for _, l := range remainingLetters {
			fmt.Print(string(l) + " ")
		}

		// Print remaining attempts
		attemptsRemaining--
		fmt.Println("\nAttempts remaining: " + fmt.Sprint(attemptsRemaining))
		
		// If there is no more attempts it means that the game is lost
		if attemptsRemaining == 0 { 
			currentStats.victory = false
			break
		}

	}

	// Storing post-game data
	currentStats.secretWord = string(wordToGuess)
	currentStats.attempts = 6 - attemptsRemaining

	saveStats()
	if wantStats(scanner) { showStats() }

	fmt.Println("Press Enter to exit...")
	scanner.Scan()
}