# 🎙️ Wuthering Waves Voice-over Downloader

A simple and lightweight Go-based tool to automatically download character voice-overs from [Encore.moe](https://encore.moe/). This project was built for personal learning and to support contributions to the [Wuthering Waves Fandom Wiki](https://wutheringwaves.fandom.com/wiki/).

---

## 🚀 Features

* 📁 Downloads voice files from Encore.moe in `.mp3` format
* 🔃 Convert the voice files from `.mp3` to `.ogg` automatically (This is optional)
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

### Run the Program

Make sure you have the latest version of [Go](https://go.dev/doc/install) installed.

Then, install depedencies and run:

```bash
go mod tidy
go run main.go
```

You will be asked which character voice you want to download. Just submit and wait for the program finished. Voice lines will be downloaded into the `voices/` directory.

#### 💡 FOR NON DEVELOPER USER 💡

You can download the executable program [here](https://github.com/Iqrar99/WW-VO-Downloader/releases) then put the `.exe` file in the repository after cloning it.

---

## 📝 Note for Fandom Contributors

If you're contributing to the [Wuthering Waves Fandom Wiki](https://wutheringwaves.fandom.com/wiki/), you’ll need to convert the `.mp3` files to `.ogg` format before uploading.

There are 2 ways to do that:

1. Download the ffmpeg static build from [official website](https://ffmpeg.org/download.html) then extract and put the file in `./engine` dir. Read the [engine/README.md](./engine/README.md) file for more info. After that, turn on the automatic conversion during program start. Or,

2. Use [AscensionGameDev's Batch MP3 to OGG Converter](https://github.com/AscensionGameDev/Batch-MP3-to-OGG-Converter) for fast and easy batch conversion. This is needed if you turn off the automatic conversion.

---

## 🪲 Found a Bug?

If you run into an issue or unexpected behavior, feel free to [open an issue](https://github.com/your-username/wuthering-waves-voice-downloader/issues) and let me know what’s happening. The more detail, the better!

---

## 🙏 Special Thanks

A big shoutout to [@alteria0](https://twitter.com/alteria0), the developer behind [Encore.moe](https://encore.moe/), for creating such an amazing resource for the community.
