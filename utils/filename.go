package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// Voice filename is following Fandom WIKI filename format
func composeVoiceFileName(resonator, lang, title string, wikiMode bool) string {
	title = handleTitle(title, wikiMode)

	var fileName string
	if lang == "En" {
		fileName = fmt.Sprintf("%s %s", resonator, title)
	} else {
		fileName = fmt.Sprintf("%s %s %s", resonator, strings.ToUpper(lang), title)
	}
	fileName = strings.ReplaceAll(fileName, ":", "")
	fileName = strings.ReplaceAll(fileName, " ", "_")

	return fileName + ".mp3"
}

// Handle specific title formats
func handleTitle(title string, wikiMode bool) string {
	if strings.Contains(title, "'s Hobby") {
		return "Hobby"
	}
	if strings.Contains(title, "'s Trouble") {
		return "Trouble"
	}
	if wikiMode {
		title = handleCombatTitle(title)
	}
	return title
}

func handleCombatTitle(title string) string {
	// Handle very special case
	if strings.Contains(title, "Intro Skill:") {
		title = strings.ReplaceAll(title, "Intro Skill:", "Intro & Outro Skills:")
	}

	// Check for sole title first
	for _, t := range soleTitles {
		if strings.EqualFold(title, t) {
			return title + " 01"
		}
	}

	// Handle non-sole title with roman numerals
	pattern := regexp.MustCompile(`(.+:\s)([IVXLCDM]+)$`)
	for _, t := range nonSoleTitles {
		if strings.Contains(title, t) {
			return pattern.ReplaceAllStringFunc(title, func(s string) string {
				matches := pattern.FindStringSubmatch(s)
				prefix := matches[1]
				num := romanToInt(matches[2])
				return prefix + fmt.Sprintf("%02d", num)
			})
		}
	}

	return title
}
