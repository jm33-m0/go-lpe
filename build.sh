#!/bin/bash

(
	echo "building cmd/golpe/main.go"
	go mod tidy &&
		cd cmd/golpe &&
		CGO_ENABLED=0 go build -o ../../golpe.exe -ldflags="-s -w" -trimpath
	command -v upx >/dev/null && upx ../../golpe.exe
)

(
	echo "building demo.c"
	musl-gcc -static -s -o emp3r0r.exe demo.c || echo "musl-gcc not found"
	command -v upx >/dev/null && upx emp3r0r.exe
	echo "Rename emp3r0r.exe to emp3r0r and use it with golpe.exe"
)
