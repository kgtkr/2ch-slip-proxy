#!/bin/sh
GOOS=windows GOARCH=amd64 go build -o bin/2ch-slip-proxy_v$1_windows_amd64.exe
GOOS=linux GOARCH=amd64 go build -o bin/2ch-slip-proxy_v$1_linux_amd64
GOOS=darwin GOARCH=amd64 go build -o bin/2ch-slip-proxy_v$1_mac_amd64