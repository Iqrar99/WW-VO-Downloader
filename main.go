package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Iqrar99/ww-vo-downloader/utils"
)

var (
	resonator string
	languages = []string{"En", "Ja", "Ko", "Zh"}
)

func main() {
	var jsonFileName string
	fmt.Print("Enter the JSON file name (without extension): ")
	fmt.Scanln(&jsonFileName)
	jsonFileName = strings.ToLower(jsonFileName)
	resonator = strings.ToUpper(jsonFileName[:1]) + jsonFileName[1:]

	jsonData := utils.ReadJsonFile(jsonFileName + ".json").(map[string]any)
	for _, lang := range languages {
		currPath := "voices/" + resonator + "/" + lang
		utils.CreateDir(currPath)
		log.Printf("Processing %s pack\n", strings.ToUpper(lang))

		for _, words := range jsonData["Words"].([]any) {
			wordMap := words.(map[string]any)

			title := wordMap["Title"].(string)
			voiceURL := wordMap["Voice"+lang].(string)
			err := utils.DownloadVoiceFile(voiceURL, currPath, resonator, lang, title)
			if err != nil {
				log.Println("Error:", err)
			} else {
				log.Printf("Saved: %s - %s - %s", resonator, strings.ToUpper(lang), title)
			}

			// Wait to avoid rate limits
			time.Sleep(2 * time.Second)
		}
	}
	log.Println("All voice files downloaded successfully.")
}
