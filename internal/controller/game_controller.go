package controller

import (
	"fmt"

	"github.com/remiposo/wordle-cli/internal/config"
	"github.com/remiposo/wordle-cli/internal/model"
	"github.com/remiposo/wordle-cli/internal/presentation"
)

type GameController struct {
}

func (gc *GameController) Play() error {
	words, err := config.LoadWords()
	if err != nil {
		return fmt.Errorf("unable to load words config: %w", err)
	}
	game := model.NewGame(words)
	for game.CanAnswer() {
		answer := presentation.ReadAnswer(game.NextAnswerCount())
		result, err := game.SubmitAnswer(answer)
		if err != nil {
			presentation.PrintError(err)
			continue
		}
		presentation.PrintResult(result)
		presentation.PrintStatusMap(game.Results.GetStatusMap())
	}
	presentation.PrintResults(game.Target, game.Results)
	return nil
}
