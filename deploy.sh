#!/bin/bash



VERSION="v1.0.1"



CGO_ENABLED=0  GOOS=linux   GOARCH=amd64  go build -o ./build/linux/cdn_refresh -v ./src/main.go
tar -czvf ./build/linux/cdn-refresh-${VERSION}-linux-amd64.tar.gz ./build/linux/cdn_refresh



CGO_ENABLED=0  GOOS=darwin  GOARCH=amd64  go build -o ./build/mac/cdn_refresh -v ./src/main.go
tar -czvf ./build/mac/cdn-refresh-${VERSION}-mac-amd64.tar.gz ./build/mac/cdn_refresh



CGO_ENABLED=0  GOOS=windows  GOARCH=amd64  go build -o ./build/windows/cdn_refresh.exe -v ./src/main.go
tar -czvf ./build/windows/cdn-refresh-${VERSION}-windows-amd64.tar.gz ./build/windows/cdn_refresh.exe


