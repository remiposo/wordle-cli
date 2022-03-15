package wordlecli

import (
	"github.com/remiposo/wordle-cli/internal/controller"
	"github.com/urfave/cli/v2"
)

func Play(c *cli.Context) error {
	gameController := new(controller.GameController)
	return gameController.Play()
}
