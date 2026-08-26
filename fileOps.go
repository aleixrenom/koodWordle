package main

import (
	"os"
	"fmt"
	"bufio"
	"encoding/csv"
	"io"
	"strconv"
)

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

func getWordToGuess(wordNumber int) []rune {
	word := []rune{}
	wordsFile, fErr := os.OpenFile("wordle-words.txt", os.O_RDONLY, 0644)
	if fErr != nil {
		fmt.Println("Error opening the words file:", fErr)
		os.Exit(0)
	}

	wordsScanner := bufio.NewScanner(wordsFile)

	lineNum := 0
	for {
		if wordsScanner.Scan() {
			lineNum++
			if lineNum - 1 == wordNumber { 
				word = []rune(wordsScanner.Text())
				break
			}
		} else {
			fmt.Println("Not enough words available. Please, give a smaller number as the first argument.")
			wordsFile.Close()
			os.Exit(0)
		}
	}
	wordsFile.Close()

	if len(word) == 0 {
		fmt.Println("There has been an error finding the word. We apologize.")
		os.Exit(0)
	}

	return word
}

func saveStats() {
	// Open stats
	statsFile, cErr := os.OpenFile("stats.csv", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0644)
	if cErr != nil {
		fmt.Println("Error opening the stats file:", cErr)
		os.Exit(0)
	}

	// Save stats
	statsWriter := csv.NewWriter(statsFile)

	if ssErr := statsWriter.Write([]string{
		currentStats.username,
		currentStats.secretWord,
		fmt.Sprint(currentStats.attempts),
		fmt.Sprint(currentStats.victory),
	}); ssErr != nil {
		fmt.Println("Error updating the stats file:", ssErr)
		os.Exit(0)
	}

	statsWriter.Flush() // Write to disk

	statsFile.Close()
}

func showStats() {
	// Open stats
	statsFile, cErr := os.OpenFile("stats.csv", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0644)
	if cErr != nil {
		fmt.Println("Error opening the stats file:", cErr)
		os.Exit(0)
	}
	
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
	fmt.Println("Average attempts per game:", strconv.FormatFloat(attemptsAvg, 'f', 2, 64))

	statsFile.Close()
}