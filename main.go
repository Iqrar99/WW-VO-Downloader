package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Iqrar99/ww-vo-downloader/utils"
)

var (
	wikiMode     bool
	autoConvert  bool
	resonator    string
	jsonFileName string
	languages    = map[string]bool{
		"En": false,
		"Ja": false,
		"Ko": false,
		"Zh": false,
	}
)

func main() {
	startInteractiveInput()

	jsonFileName = strings.ToLower(jsonFileName)
	resonator = strings.ToUpper(jsonFileName[:1]) + jsonFileName[1:]

	jsonData := utils.ReadJsonFile(jsonFileName + ".json").(map[string]any)
	for lang := range languages {
		if !languages[lang] {
			continue
		}

		currPath := "voices/" + resonator + "/" + lang
		utils.CreateDir(currPath)
		if wikiMode {
			log.Printf("Processing %s pack with WIKI mode on\n", strings.ToUpper(lang))
		} else {
			log.Printf("Processing %s pack\n", strings.ToUpper(lang))
		}

		for _, words := range jsonData["Words"].([]any) {
			wordMap := words.(map[string]any)

			title := wordMap["Title"].(string)
			voiceURL := wordMap["Voice"+lang].(string)
			err := utils.DownloadVoiceFile(voiceURL, currPath, resonator, lang, title, wikiMode)
			if err != nil {
				log.Println("Error:", err)
			} else {
				log.Printf("Saved: %s - %s - %s", resonator, strings.ToUpper(lang), title)
			}

			// Wait to avoid rate limits
			time.Sleep(2 * time.Second)
		}

		log.Printf(
			"All %s voice files have downloaded successfully.",
			strings.ToUpper(lang),
		)

		if autoConvert {
			log.Printf(
				"Processing %s voice files for conversion to OGG format...",
				strings.ToUpper(lang),
			)
			utils.ConvertVoiceFiles(currPath)
			log.Println("Conversion completed successfully.")
		}
	}
	log.Println("All voice files have processed thoroughly.")
	time.Sleep(5 * time.Second)
}

func startInteractiveInput() {
	var userChoice string
	var isDownloadAll bool

	fmt.Print("Enter the JSON file name (without extension): ")
	fmt.Scanln(&jsonFileName)
	utils.HandleEmptyInput(jsonFileName)

	utils.PrintSeparator()

	fmt.Println("Turn WIKI mode on? (y/n)")
	fmt.Print("By turning on this mode, certain files will be renamed to adjust Fandom WIKI format: ")
	fmt.Scanln(&userChoice)
	utils.HandleEmptyInput(userChoice)
	utils.HandleYesNoInput(userChoice, &wikiMode)

	utils.PrintSeparator()

	fmt.Print("Do you want to download all voices from all languages? (y/n): ")
	fmt.Scanln(&userChoice)
	utils.HandleEmptyInput(userChoice)
	utils.HandleYesNoInput(userChoice, &isDownloadAll)
	if isDownloadAll {
		for lang := range languages {
			languages[lang] = true
		}
		return
	}

	utils.PrintSeparator()

	fmt.Print("Which language pack you want to download? (EN, JA, KO, ZH): ")
	fmt.Scanln(&userChoice)
	switch strings.ToUpper(userChoice) {
	case "EN":
		languages["En"] = true
	case "JA":
		languages["Ja"] = true
	case "KO":
		languages["Ko"] = true
	case "ZH":
		languages["Zh"] = true
	default:
		log.Fatal("Invalid language choice. Please enter EN, JA, KO, or ZH.")
	}

	utils.PrintSeparator()

	fmt.Println("Turn on Automatic Conversion? (y/n)")
	fmt.Print("By turning on this mode, files will be converted from mp3 to ogg: ")
	fmt.Scanln(&userChoice)
	utils.HandleEmptyInput(userChoice)
	utils.HandleYesNoInput(userChoice, &autoConvert)
	if autoConvert {
		utils.CheckFFmpegExists()
	}

	utils.PrintSeparator()
}
