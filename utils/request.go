package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	fakeUA "github.com/lib4u/fake-useragent"
)

var client = &http.Client{Timeout: 30 * time.Second}

func DownloadCharacterData(characterName string) error {
	_, err := os.Stat("./json/" + characterName + ".json")
	if !os.IsNotExist(err) {
		return nil
	}
	log.Printf("Proceed to download character [%s] JSON data...", characterName)

	url := fmt.Sprintf("https://api.encore.moe/en/character/%s.json", CharacterData[characterName])
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download: %s (%d)", url, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var prettyJSON any
	if err := json.Unmarshal(body, &prettyJSON); err != nil {
		return err
	}

	prettyBody, err := json.MarshalIndent(prettyJSON, "", "  ")
	if err != nil {
		return err
	}

	out, err := os.Create(fmt.Sprintf("json/%s.json", characterName))
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := out.Write(prettyBody); err != nil {
		return err
	}

	log.Printf("Data %s.json has downloaded successfully", characterName)
	return nil
}

func DownloadVoiceFile(url, path, resonator, lang, title string, wikiMode bool) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	err = addCustomHeaders(req)
	if err != nil {
		return err
	}

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

func addCustomHeaders(req *http.Request) error {
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

	return nil
}
