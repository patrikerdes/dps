#!/usr/bin/env bash

echo "cross compiling..."
mkdir -p osx linux windows

set -x
env GOOS=darwin GOARCH=amd64 go build -o osx/dps2mpd dps2mpd.go
env GOOS=darwin GOARCH=amd64 go build -o osx/mpd2dps mpd2dps.go

env GOOS=windows GOARCH=amd64 go build -o windows/dps2mpd.exe dps2mpd.go
env GOOS=windows GOARCH=amd64 go build -o windows/mpd2dps.exe mpd2dps.go

env GOOS=linux GOARCH=amd64 go build -o linux/dps2mpd dps2mpd.go
env GOOS=linux GOARCH=amd64 go build -o linux/mpd2dps mpd2dps.go
