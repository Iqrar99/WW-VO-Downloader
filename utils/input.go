package utils

import (
	"log"
	"strings"

	"github.com/xrash/smetrics"
)

func HandleEmptyInput(input string) {
	if input == "" {
		log.Fatal("input can't be empty.")
	}
}

func HandleYesNoInput(input string, target *bool) {
	if strings.ToLower(input) == "y" || strings.ToLower(input) == "yes" {
		*target = true
	} else if strings.ToLower(input) == "n" || strings.ToLower(input) == "no" {
		*target = false
	} else {
		log.Fatal("Invalid input. Please enter 'y' or 'n'.")
	}
}

func HandleCharacterInput(characterName *string) {
	var similarName, name string
	var score float64
	maxScore := -1.0
	for n := range CharacterData {
		name = strings.ReplaceAll(n, "_", " ")
		score = smetrics.JaroWinkler(*characterName, name, 0.7, 4)
		if score >= 0.7 && score > maxScore {
			similarName = name
			maxScore = score
		}
	}
	if similarName == "" {
		log.Fatal("Can't find character name.")
	}
	*characterName = similarName
}
