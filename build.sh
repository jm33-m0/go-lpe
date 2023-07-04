#!/bin/bash

(
	echo "building cmd/golpe/main.go"
	go mod tidy &&
		cd cmd/golpe &&
		CGO_ENABLED=0 go build -o ../../golpe.exe -ldflags="-s -w" -trimpath
	command -v upx >/dev/null && upx ../../golpe.exe
)

(
	echo "building cmd/demo/main.go"
	cd cmd/demo &&
		CGO_ENABLED=0 go build -o ../../emp3r0r.exe -ldflags="-s -w" -trimpath
	command -v upx >/dev/null && upx ../../emp3r0r.exe
	echo "Copy ./emp3r0r.exe to your target environment and rename it to emp3r0r. Test ./golpe with it, you should get a root shell"
)
