# 🎙️ Wuthering Waves Voice-over Downloader

A simple and lightweight Go-based tool to automatically download character voice-overs from [Encore.moe](https://encore.moe/). This project was built for personal learning and to support contributions to the [Wuthering Waves Fandom Wiki](https://wutheringwaves.fandom.com/wiki/).

---

## 🚀 Features

* 📁 Downloads voice files from Encore.moe in `.mp3` format
* 🛠️ Lightweight and written in pure Go
* 🚧 Built-in rate limiting to avoid HTTP 429 (Too Many Requests) errors
* 🎓 Great for learning or contributing to fan-driven projects

---

## 📦 Getting Started

Follow the steps below to get everything set up:

> 💡 Clone the repo first using:

```bash
git clone https://github.com/Iqrar99/WW-VO-Downloader.git
cd WW-VO-Downloader
```

### 1. Find the Character

* Visit the [Encore.moe character list](https://encore.moe/character).
* Click on a character you’re interested in.
* Take note of the numeric ID from the URL.

  * *Example:* Camellya’s page is `https://encore.moe/character/1603`, so the ID is `1603`.

### 2. Download the JSON File

* Use the API endpoint:
  `https://api.encore.moe/en/character/<id>.json`
  Replace `<id>` with the actual character ID.
* Save the file in the `json/` directory.
* Rename it to something meaningful, like:
  `1603.json → camellya.json`

### 3. Run the Program

Make sure you have the latest version of [Go](https://go.dev/doc/install) installed.

Then, run:

```bash
go run main.go
```

You will be asked which `.json` file you want to use. Just submit and wait for the program finished. Voice lines will be downloaded into the `voices/` directory.

---

## 📝 Note for Fandom Contributors

If you're contributing to the [Wuthering Waves Fandom Wiki](https://wutheringwaves.fandom.com/wiki/), you’ll need to convert the `.mp3` files to `.ogg` format before uploading.

I recommend using [AscensionGameDev's Batch MP3 to OGG Converter](https://github.com/AscensionGameDev/Batch-MP3-to-OGG-Converter) for fast and easy batch conversion.

---

## 🪲 Found a Bug?

If you run into an issue or unexpected behavior, feel free to [open an issue](https://github.com/your-username/wuthering-waves-voice-downloader/issues) and let me know what’s happening. The more detail, the better!

---

## 🙏 Special Thanks

A big shoutout to [@alteria0](https://twitter.com/alteria0), the developer behind [Encore.moe](https://encore.moe/), for creating such an amazing resource for the community.
