
default: watch

alias b := build
alias t := tidy

[parallel]
watch: server-watch open-firefox

server-watch:
    gin run main.go --bin gin-bin
    
open-firefox:
    firefox --new-window "http://localhost:3000"

build:
    go build main.go

tidy:
    go mod tidy

clean:
    go clean