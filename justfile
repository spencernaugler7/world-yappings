[windows]
set shell := ["powershell.exe", "-NoLogo", "-Command"]

[unix]
set shell := ["sh", "-cu"]

alias b := build
alias t := tidy

default: watch

watch:
    air server --port 8080

build:
    go build main.go

tidy:
    go mod tidy

clean:
    go clean

install-deps:
    go install github.com/air-verse/air@latest
