# go-lpe
A collection of LPE exploits written in Go

## Exploits

| CVE        | Description                                                                        | Link                                          |
| 2021-4034  | pkexec, using `execve` syscall in Go doesn't work, so we just run blasty's exploit | https://haxx.in/files/blasty-vs-pkexec.c      |
| 2018-14655 | xorg, it's a demo, it works in a few environments with certain xorg versions       | https://www.cvedetails.com/cve/CVE-2018-14655 |

## Usage

```bash
go get -u -v github.com/jm33-m0/go-lpe
```

```go
package main

import (
    gople "github.com/jm33-m0/go-lpe"
)

func main() {
    gople.RunAll()
}
```

Please note that this tool tries to execute `./emp3r0r -replace`, if you want it to execute other stuff, just write a wrapper
