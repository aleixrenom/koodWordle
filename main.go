package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"unicode"
	"encoding/csv"
	"io"
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

	fmt.Printf("(Don't tell anyone, but the word is \"%s\")\n", string(wordToGuess))

	// --------------------------------------------------------------
	// 
	// --------------------------------------------------------------

	currentStats := GameData{}
	attemptsRemaining := 6
	gameFinished := false

	scanner := bufio.NewScanner(os.Stdin)

	for currentStats.username == "" {
		fmt.Print("Enter your username: ")
		if scanner.Scan() {
			currentStats.username = strings.TrimSpace(scanner.Text())
		}
	}

	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")
	
	for {
		guess := ""
		fmt.Print("Enter your guess: ")
		if scanner.Scan() {
			guess = strings.TrimSpace(scanner.Text())
		} else {
			fmt.Printf("There was a problem gathering input.")
			break
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
				break
			}
		}
		if con { continue }

		// Correct guess
		if guess == string(wordToGuess) {
			fmt.Println("Congratulations! You've guessed the word correctly.")
			currentStats.victory = true
			gameFinished = true
			break
		}

		// Guessed word processing
		fmt.Print("Feedback: ")
		for i, ch := range guess {
			uppCh := ch - 'a' + 'A'
			if ch == wordToGuess[i] {
				fmt.Print(Green + string(uppCh) + Reset)
				continue
			} else if strings.ContainsRune(string(wordToGuess), ch) {
				fmt.Print(Yellow + string(uppCh) + Reset)
				continue
			} else {
				fmt.Print(White + string(uppCh) + Reset)
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

		// Remaining letters
		fmt.Print("\nRemaining letters: ")
		for _, l := range remainingLetters {
			fmt.Print(string(l) + " ")
		}

		// Remaining attempts
		attemptsRemaining--
		fmt.Println("\nAttempts remaining: " + fmt.Sprint(attemptsRemaining))
		
		if attemptsRemaining == 0 { 
			currentStats.victory = false
			gameFinished = true
			break 
		}

	}

	// Storing post-game data
	currentStats.secretWord = string(wordToGuess)
	currentStats.attempts = 6 - attemptsRemaining

	// In some cases where validation has not passed, the execution will arrive here
	// with gameFinished still being false
	if !gameFinished { return }

	// Ask for stats
	wantStats := ""
	fmt.Print("Do you want to see your stats? (yes/no): ")
	for wantStats != "yes" && wantStats != "no" {
		if scanner.Scan() {
			wantStats = strings.TrimSpace(scanner.Text())
		} else {
			fmt.Println("\nThere was an error processing your input.")
			break
		}
		if wantStats != "yes" && wantStats != "no" {
			fmt.Print("Please, wright only 'yes' or 'no': ")
		}

	}

	// Open stats
	statsFile, cErr := os.OpenFile("stats.csv", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0644)
	if cErr != nil {
		fmt.Println("Error opening the stats file:", cErr)
		return
	}
	defer statsFile.Close() // the program will end very soon, so we can trust that the defer will trigger soon too

	// Save stats
	statsWriter := csv.NewWriter(statsFile)

	if ssErr := statsWriter.Write([]string{
		currentStats.username,
		currentStats.secretWord,
		fmt.Sprint(currentStats.attempts),
		fmt.Sprint(currentStats.victory),
	}); ssErr != nil {
		fmt.Println("Error updating the stats file:", ssErr)
		return
	}

	statsWriter.Flush() // Write to disk

	if wantStats == "yes" {
		// Show stats
		statsFile.Seek(0,0) // After saving the stats into the file, the cursor is at the end, so we have to bring it back up to read
		statsReader := csv.NewReader(statsFile)
		userStats := [][]string{}

		for {
			record, srErr := statsReader.Read()
			if srErr == io.EOF {
				break
			}
			if srErr != nil {
				fmt.Println("Error reading the stats file:", srErr)
				break
			}

			if currentStats.username == record[0] {
				userStats = append(userStats, record)
			}
		}

		gamesWon := 0
		attemptsSum := 0

		for _, match := range userStats {
			if match[3] == "true" { gamesWon++ }
			attemptsAmt, _ := strconv.Atoi(match[2])
			attemptsSum += attemptsAmt
		}

		attemptsAvg := float64(attemptsSum) / float64(len(userStats)) // 64 because strconv.FormatFloat() asks for 64 (I think)

		fmt.Println("Stats for " + userStats[0][0] + ":")
		fmt.Println("Games played:", len(userStats))
		fmt.Println("Games won:", gamesWon)
		fmt.Println("Average attempts per game:", strconv.FormatFloat(attemptsAvg, 'f', 1, 64))
	}
}

func wordExists(word string) bool {
	found := false
	wordsFile, oErr := os.OpenFile("wordle-words.txt", os.O_RDONLY, 0644)
	if oErr != nil {
		fmt.Println("Error opening the words file:", oErr)
		return found
	}

	wordsScanner := bufio.NewScanner(wordsFile)

	for {
		if wordsScanner.Scan() {
			if word == wordsScanner.Text() { 
				found = true
				break
			}
		} else { // EOF
			break
		}
	}
	wordsFile.Close()
	return found
}