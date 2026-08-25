package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"unicode"
)

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

func main() {
	args := os.Args[1:] // grab all user-given arguments (first one is always just path)
	// scanner := bufio.NewScanner(os.Stdin)

	if len(args) != 1 {
		fmt.Println("Usage: go run . [number]")
		return
	}

	// --------------------------------------------------------------
	// GET WORD TO GUESS
	// --------------------------------------------------------------
	wordToGuess := []rune{}

	wordNumber, wErr := strconv.Atoi(args[0])
	if wErr != nil {
		fmt.Println("Usage: go run . [number]\n", wErr)
		return
	}

	if wordNumber <= 0 {
		fmt.Println("Please, give a positive number as the first argument.")
		return
	}

	wordsFile, fErr := os.OpenFile("wordle-words.txt", os.O_RDONLY, 0644)
	if fErr != nil {
		fmt.Println("Error opening the words file:", fErr)
		return
	}

	wordsScanner := bufio.NewScanner(wordsFile)

	lineNum := 0
	for {
		if wordsScanner.Scan() {
			lineNum++
			if lineNum == wordNumber { 
				wordToGuess = []rune(wordsScanner.Text())
				break
			}
		} else {
			fmt.Println("Not enough words available. Please, give a smaller number as the first argument.")
			wordsFile.Close()
			return
		}
	}
	wordsFile.Close()

	if len(wordToGuess) == 0 {
		fmt.Println("There has been an error finding the word. We apologize.")
		return
	}

	// --------------------------------------------------------------
	// 
	// --------------------------------------------------------------

	currentStats := GameData{}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your username: ")
	if scanner.Scan() {
		currentStats.username = strings.TrimSpace(scanner.Text())
	}

	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")
	
	for scanner.Scan() {
		fmt.Print("Enter your guess: ")
		guess := strings.TrimSpace(scanner.Text())

		// Input validation
		if len(guess) != 5 {
			fmt.Println("Your guess must be exactly 5 letters long.")
			continue
		}
		con := false
		for _, ch := range guess {
			if unicode.IsDigit(ch) || unicode.IsUpper(ch) {
				fmt.Println("Your guess must only contain lowercase letters.")
				con = true
				break
			}
		}
		if con { continue }

		// Guessed word processing
		fmt.Print("Feedback: ")
		for i, ch := range guess {
			uppCh := ch - 'a' + 'A'
			if ch == wordToGuess[i] {
				fmt.Print(Green + string(uppCh) + Reset + " ")
				continue
			} else if strings.ContainsRune(string(wordToGuess), ch) {
				fmt.Print(Yellow + string(uppCh) + Reset + " ")
				continue
			} else {
				fmt.Print(White + string(uppCh) + Reset + " ")
				// remove it from the remaining letters
				for j, l := range remainingLetters {
					if uppCh == l {
						if len(remainingLetters) == 1 {
							remainingLetters = []rune{}
						} else if j == 0 {
							remainingLetters = remainingLetters[1:]
						} else {
							remainingLetters = append(remainingLetters[:j], remainingLetters[j+1:]...)
						}
						break
					}
				}
			}
		}
	}

}

// User enters the username they want to use
	// This username is stored during execution

// User can start guessing with 6 attempts
	// Validate the input

// After validation, display the feedback and show the remaining letters

// After the guesses run out or they guess the word, ask them if they want to see stats

// Regardless if they displayed stats or not, do the "press enter to exit"

/*
What is the deal with the database?
It stores: username, word to guess, attempts used, win/lose
*/
