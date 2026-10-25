[windows]
set shell := ["powershell.exe", "-NoLogo", "-Command"]

[unix]
set shell := ["sh", "-cu"]

alias b := build
alias t := tidy

default: watch

[parallel]
watch: server-watch open-firefox

server-watch:
    air server --port 8080
    
open-firefox:
    firefox --new-window "http://localhost:8080"

build:
    go build main.go

tidy:
    go mod tidy

clean:
    go clean

install-deps:
    go install github.com/air-verse/air@latest