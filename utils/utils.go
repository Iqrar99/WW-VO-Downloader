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
)

var client = &http.Client{
	Timeout: 30 * time.Second,
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
func composeVoiceFileName(resonator, lang, title string) string {
	title = handleTitle(title)

	var fileName string
	if lang == "En" {
		fileName = fmt.Sprintf("%s %s", resonator, title)
	} else {
		fileName = fmt.Sprintf("%s %s %s", resonator, strings.ToUpper(lang), title)
	}
	fileName = strings.Replace(fileName, ":", "", -1)
	fileName = strings.Replace(fileName, " ", "_", -1)

	return fileName + ".mp3"
}

// Handle specific title formats
func handleTitle(title string) string {
	matched, _ := regexp.MatchString(`.+\'s Hobby`, title)
	if matched {
		return "Hobby"
	}

	matched, _ = regexp.MatchString(`.+\'s Trouble`, title)
	if matched {
		return "Trouble"
	}

	return title
}

func DownloadVoiceFile(url, path, resonator, lang, title string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Add headers to look like a real browser
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_8_5) Gecko/20100101 Firefox/57.3")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle rate limiting (429)
	if resp.StatusCode == http.StatusTooManyRequests {
		log.Println("Rate limited. Waiting before retry...")
		time.Sleep(10 * time.Second)
		return DownloadVoiceFile(url, path, resonator, lang, title)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download: %s (%d)", url, resp.StatusCode)
	}

	out, err := os.Create(path + "/" + composeVoiceFileName(resonator, lang, title))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
