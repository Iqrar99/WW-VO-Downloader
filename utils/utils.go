package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var (
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
	ffmpegPath    string
	CharacterData map[string]string
)

func init() {
	if runtime.GOOS == "windows" {
		ffmpegPath = "./engine/ffmpeg.exe"
	} else {
		ffmpegPath = "./engine/ffmpeg"
	}

	rawData := ReadJsonFile("data/", "character.json")
	CharacterData = make(map[string]string)
	for name, v := range rawData {
		if idVal, ok := v.(map[string]any)["id"].(string); ok {
			CharacterData[name] = idVal
		}
	}
}

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

func PrintSeparator() {
	fmt.Println(strings.Repeat("-", 50))
}

func ReadJsonFile(path, jsonFileName string) map[string]any {
	jsonFile, err := os.Open(path + jsonFileName)
	if err != nil {
		log.Fatal("Error opening JSON file:", err)
	}
	defer jsonFile.Close()

	var data map[string]any
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

// Checks if the ffmpeg program exists based on OS
func CheckFFmpegExists() {
	_, err := os.Stat(ffmpegPath)
	if os.IsNotExist(err) {
		log.Fatal(
			"FFmpeg program not found in ./engine directory. " +
				"Consider downloading it from https://ffmpeg.org/download.html " +
				"and placing it in the engine directory.",
		)
	}
}

func ConvertVoiceFiles(path string) {
	var inputFile string
	var outputFile string
	var cmd *exec.Cmd

	entries, _ := os.ReadDir(path)
	for _, entry := range entries {
		inputFile = filepath.Join(path, entry.Name())
		outputFile = strings.ReplaceAll(inputFile, ".mp3", ".ogg")

		cmd = exec.Command(
			ffmpegPath,
			"-y",
			"-i", inputFile,
			"-c:a", "libvorbis",
			outputFile,
		)
		err := cmd.Run()
		if err != nil {
			log.Printf("Error converting %s: %v", inputFile, err)
			continue
		}
		os.Remove(inputFile)

		log.Println("Converted to OGG:", entry.Name()[:len(entry.Name())-4])
	}
}
