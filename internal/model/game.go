package model

import (
	"errors"
	"math/rand"
	"time"
	"unicode/utf8"
)

type Game struct {
	MaxAnswerCount int
	Results        Results
	Target         string
	Words          Words
}

const (
	MAX_ANSWER_COUNT = 6
)

var (
	ErrInvalidLength = errors.New("invalid word length")
	ErrWordNotFound  = errors.New("word not found")
)

func NewGame(words Words) *Game {
	game := &Game{
		Words:          words,
		MaxAnswerCount: MAX_ANSWER_COUNT,
	}
	game.initTarget()
	return game
}

func (g *Game) initTarget() {
	rand.Seed(time.Now().UnixNano())
	g.Target = g.Words[rand.Intn(len(g.Words))]
}

func (g *Game) CanAnswer() bool {
	return !g.Solved() &&
		g.NextAnswerCount() <= g.MaxAnswerCount
}

func (g *Game) Solved() bool {
	return g.Results.Solved()
}

func (g *Game) NextAnswerCount() int {
	return len(g.Results) + 1
}

func (g *Game) SubmitAnswer(answer string) (*Result, error) {
	if utf8.RuneCountInString(answer) != g.Words.WordLength() {
		return nil, ErrInvalidLength
	}
	if !g.Words.Contain(answer) {
		return nil, ErrWordNotFound
	}
	result := NewResult(answer, g.Target)
	g.Results = append(g.Results, result)
	return result, nil
}
