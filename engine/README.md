# Note

This folder is intended to contain `ffmpeg` binary file or executable file (depends on your OS).

Download the static build from [official website](https://ffmpeg.org/download.html). Then extract and put the `bin` or `.exe` file in this folder.

It will be like this:

```shell
this-project/
├── engine/
│   ├── ffmpeg      <- static binary ffmpeg (If you use Linux/Ubuntu)
│   ├── ffmpeg.exe  <- static build ffmpeg (If you use Windows instead)
│   └── README.md
├── main.go
...
```
