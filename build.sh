#!/bin/bash

echo "$0 builds cmd/golpe/main.go"
(
    cd cmd/golpe &&
        CGO_ENABLED=0 go build -o ../../golpe.exe -ldflags="-s -w" -trimpath &&
        command -v upx && upx ../../golpe.exe
)
