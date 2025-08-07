# 🧾 Expected JSON Structure

Here’s a simplified look at the structure your `.json` file should follow:

```json
{
    // ...
    "Words": [
        {
            "Id": 140701,
            "RoleId": 1407,
            "Type": 1,
            "Sort": 1,
            "Title": "Thoughts: I",
            "Content": "Some dialogue text here...",
            "Voice": "/Game/ExamplePath/play_example.play_example",
            "VoiceEn": "https://api.encore.moe/resource/ExamplePath/en/example.mp3",
            "VoiceJa": "https://api.encore.moe/resource/ExamplePath/ja/example.mp3",
            "VoiceKo": "https://api.encore.moe/resource/ExamplePath/ko/example.mp3",
            "VoiceZh": "https://api.encore.moe/resource/ExamplePath/zh/example.mp3"
        }
    ]
    // ...
}
```

## 🎯 What This Program Actually Uses

From the full JSON, the downloader only uses these specific fields:

* `Title` — Used for naming and reference
* `VoiceEn` — English voice line URL
* `VoiceJa` — Japanese voice line URL
* `VoiceKo` — Korean voice line URL
* `VoiceZh` — Chinese voice line URL
