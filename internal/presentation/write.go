package presentation

import (
	"fmt"

	"github.com/remiposo/wordle-cli/internal/model"
)

func PrintError(err error) {
	fmt.Printf("%v\n\n", err)
}

func PrintLetter(letter rune, status model.Status) {
	switch status {
	case model.UNCHECKED:
		fmt.Printf("%c ", letter)
	case model.UNUSED:
		fmt.Printf("\x1b[41m%c\x1b[0m ", letter)
	case model.BITE:
		fmt.Printf("\x1b[43m%c\x1b[0m ", letter)
	case model.EAT:
		fmt.Printf("\x1b[42m%c\x1b[0m ", letter)
	}
}

func PrintResult(result *model.Result) {
	for idx, letter := range result.Answer {
		PrintLetter(letter, result.Statuses[idx])
	}
	fmt.Printf("\n\n")
}

func PrintStatusMap(statusMap map[rune]model.Status) {
	for i := 0; i < 26; i++ {
		if i > 0 && i%10 == 0 {
			fmt.Printf("\n")
		}
		letter := rune('a' + i)
		PrintLetter(letter, statusMap[letter])
	}
	fmt.Printf("\n\n")
}

func PrintResults(target string, results model.Results) {
	fmt.Printf("correct answer: %v\n", target)
	for _, result := range results {
		for idx, letter := range result.Answer {
			PrintLetter(letter, result.Statuses[idx])
		}
		fmt.Printf("\n")
	}
}
