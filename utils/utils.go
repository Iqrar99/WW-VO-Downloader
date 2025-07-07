package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	fakeUA "github.com/lib4u/fake-useragent"
)

var (
	client = &http.Client{
		Timeout: 30 * time.Second,
	}
	roman = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	soleTitles = []string{
		"Echo Summon",
		"Echo Transform",
		"Intro & Outro Skills",
		"Enemies Near",
		"Glider",
		"Grapple",
		"Sensor",
		"Wall Dash",
		"Dash",
	}
	nonSoleTitles = []string{
		"Aerial Attack:",
		"Basic Attack:",
		"Heavy Attack:",
		"Resonance Skill:",
		"Resonance Liberation:",
		"Intro & Outro Skills:",
		"Hit:",
		"Injured:",
		"Fallen:",
		"Supply Chest:",
		"Echo Summon:",
		"Echo Transform:",
		"Enemies Near:",
	}
)

// Convert roman numerals into decimal integer
func romanToInt(s string) int {
	total := 0
	prev := 0
	for i := len(s) - 1; i >= 0; i-- {
		curr := roman[s[i]]
		if curr < prev {
			total -= curr
		} else {
			total += curr
		}
		prev = curr
	}

	return total
}

func ReadJsonFile(jsonFileName string) any {
	jsonFile, err := os.Open("json/" + jsonFileName)
	if err != nil {
		log.Fatal("Error opening JSON file:", err)
	}
	log.Println("JSON file opened successfully:", jsonFileName)
	defer jsonFile.Close()

	var data any
	err = json.NewDecoder(jsonFile).Decode(&data)
	if err != nil {
		log.Fatal("Error decoding JSON file:", err)
	}

	return data
}

func CreateDir(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatalf("Error creating %s directory: %v", path, err)
	}
}

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

func DownloadVoiceFile(url, path, resonator, lang, title string, wikiMode bool) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	ua, err := fakeUA.New()
	if err != nil {
		return err
	}

	// Add headers to look like a real browser
	req.Header.Set("User-Agent", ua.GetRandom())
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://www.google.com/")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle rate limiting (429)
	if resp.StatusCode == http.StatusTooManyRequests {
		log.Println("Rate limited. Waiting before retry...")
		time.Sleep(10 * time.Second)
		return DownloadVoiceFile(url, path, resonator, lang, title, wikiMode)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download: %s (%d)", url, resp.StatusCode)
	}

	out, err := os.Create(path + "/" + composeVoiceFileName(resonator, lang, title, wikiMode))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
