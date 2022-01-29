# go-lpe
A collection of LPE exploits written in Go

![image](https://user-images.githubusercontent.com/10167884/151522601-dc61ba4b-1144-4c57-a548-f9bddd17b96e.png)

## Exploits

| CVE        | Description                                                                  | Link                                          |
|------------|------------------------------------------------------------------------------|-----------------------------------------------|
| 2021-4034  | pkexec exploit, rewritten in pure Go, base on blasty's poc                   | https://haxx.in/files/blasty-vs-pkexec.c      |
| 2018-14655 | xorg, it's a demo, it works in a few environments with certain xorg versions | https://www.cvedetails.com/cve/CVE-2018-14655 |

## Usage

```bash
go get -u -v github.com/jm33-m0/go-lpe
```

```go
package main

import (
    golpe "github.com/jm33-m0/go-lpe"
)

func main() {
    golpe.RunAll()
}
```

Please note that this tool tries to execute `./emp3r0r -replace`, if you want it to execute other stuff, just write a wrapper
