package utils

import (
	"log"
	"strings"
)

func HandleEmptyInput(input string) {
	if input == "" {
		log.Fatal("input can't be empty")
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
