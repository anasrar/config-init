#!/bin/sh

# Linux
GOOS=linux GOARCH=amd64 go build --ldflags "-s -w" -o bin/config-init cmd/config-init/main.go

# Windows
GOOS=windows GOARCH=amd64 go build --ldflags "-s -w" -o bin/config-init.exe cmd/config-init/main.go
