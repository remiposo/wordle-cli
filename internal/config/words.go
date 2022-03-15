package config

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/remiposo/wordle-cli/internal/model"
	"gopkg.in/yaml.v2"
)

type WordsConfig struct {
	Words []string `yaml:"words"`
}

func LoadWords() (model.Words, error) {
	data, err := ioutil.ReadFile(path.Join("config", "data.yaml.d", "words.yaml"))
	if err != nil {
		return nil, fmt.Errorf("unable to load file: %w", err)
	}
	wordsConfig := new(WordsConfig)
	if err := yaml.Unmarshal(data, wordsConfig); err != nil {
		return nil, fmt.Errorf("unable to parse yaml: %w", err)
	}
	words, err := model.NewWords(wordsConfig.Words)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize words instance: %w", err)
	}
	return words, nil
}
