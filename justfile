set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]
alias b := build
alias t := tidy

default: watch

[parallel]
watch: server-watch open-firefox

server-watch:
    air server --port 3000
    
open-firefox:
    firefox --new-window "http://localhost:3000"

build:
    go build main.go

tidy:
    go mod tidy

clean:
    go clean

install-deps:
    go install github.com/air-verse/air@latest