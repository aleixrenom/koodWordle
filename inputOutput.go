package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func printFeedback(guess string) {
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
}

func wantStats(scanner *bufio.Scanner) bool {
	wantStats := ""
	fmt.Print("Do you want to see your stats? (yes/no): ")
	for wantStats != "yes" && wantStats != "no" {
		if scanner.Scan() {
			wantStats = strings.TrimSpace(scanner.Text())
		} else {
			fmt.Println("\nThere was an error processing your input.")
			os.Exit(0)
		}
		if wantStats != "yes" && wantStats != "no" {
			fmt.Print("Please, wright only 'yes' or 'no': ")
		}
	}
	return wantStats == "yes"
}