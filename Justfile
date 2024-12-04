# Version information
version := `git describe --tags --always --dirty`
commit := `git rev-parse --short HEAD`
date := `date -u +"%Y-%m-%dT%H:%M:%SZ"`
ldflags := '--ldflags="-X main.version=' + version + ' -X main.commit=' + commit + ' -X main.date=' + date + '"'

default:
    @just --list

build:
    go build {{ldflags}}

test:
    go test -v ./...

test-race:
    go test -v -race ./...

test-coverage:
    go test -v -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out -o coverage.html

linux:
    GOOS=linux GOARCH=amd64 go build {{ldflags}}

clean:
    rm -f yamlsubst coverage.out coverage.html

all: test build linux

check: test build
