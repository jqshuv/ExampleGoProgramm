![Discord](https://img.shields.io/discord/903750807957147718?logo=discord&logoColor=white&style=for-the-badge)
![GitHub](https://img.shields.io/github/license/jqshuv/example-go-app?logo=github&style=for-the-badge)
![GitHub Sponsors](https://img.shields.io/github/sponsors/jqshuv?logo=github-sponsors&logoColor=white&style=for-the-badge)


# Example Go Program

This is an example program written in Go. This only opens a simple dialog box with a message in it.

## Getting Started
To get started you need the Go programming language. You can get it [here](https://go.dev). After installing it you can simply follow the next steps.

## Build
This is an example build command which uses custom flags to hide the console on Windows.

```bash 
go build -ldflags -H=windowsgui -o example.exe main.go
```

## Run
Simply run the made executable called: `example.exe`.