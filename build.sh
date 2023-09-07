#!/bin/bash

(
	echo "building cmd/golpe/main.go"
	go mod tidy &&
		cd cmd/golpe &&
		CGO_ENABLED=0 go build -o ../../golpe.exe -ldflags="-s -w" -trimpath
	command -v upx >/dev/null && upx ../../golpe.exe
)
