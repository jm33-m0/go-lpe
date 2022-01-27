#!/bin/bash

echo "$0 builds cmd/runall/main.go"
(
    cd cmd/runall &&
        CGO_ENABLED=0 go build -o ../../runall.exe -ldflags="-s -w" -trimpath &&
        command -v upx && upx ../../runall.exe
)
