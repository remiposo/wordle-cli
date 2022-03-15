package model

import (
	"fmt"
	"unicode/utf8"
)

type Words []string

func NewWords(words []string) (Words, error) {
	if err := checkLength(words); err != nil {
		return nil, err
	}
	if err := checkWordLength(words); err != nil {
		return nil, err
	}
	return Words(words), nil
}

func checkLength(words []string) error {
	if len(words) == 0 {
		return fmt.Errorf("words must not be empty")
	}
	return nil
}

func checkWordLength(words []string) error {
	length := utf8.RuneCountInString(words[0])
	for _, word := range words {
		if length != utf8.RuneCountInString(word) {
			return fmt.Errorf("words must be the same length")
		}
	}
	return nil
}

func (ws Words) WordLength() int {
	return utf8.RuneCountInString(ws[0])
}

func (ws Words) Contain(word string) bool {
	for _, w := range ws {
		if w == word {
			return true
		}
	}
	return false
}
