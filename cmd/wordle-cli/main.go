package main

import (
	"log"
	"os"

	wordlecli "github.com/remiposo/wordle-cli/internal/cmd/wordle-cli"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "wordle-cli",
		Usage:  "CLI application for playing wordle",
		Action: wordlecli.Play,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
