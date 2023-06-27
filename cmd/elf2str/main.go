package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	golpe "github.com/jm33-m0/go-lpe"
)

func main() {
	elf_path := flag.String("file", "", "target file to encode")
	flag.Parse()

	elf_data, err := ioutil.ReadFile(*elf_path)
	if err != nil {
		log.Fatalf("read %s: %v", *elf_path, err)
	}

	elf_str := golpe.Bin2String(elf_data)

	fmt.Printf("\n%s\n\n", elf_str)
	// fmt.Printf("%s\n", format_elf_str(elf_str))
}

func format_elf_str(str string) (ret string) {
	linelen := 50
	if len(str) < linelen {
		return str
	}

	// initialize
	ret = fmt.Sprintf("const elf_str = `%s`", str[:linelen])

	line := ""
	for n, c := range str {
		line += string(c)
		if n%linelen == 0 {
			ret = fmt.Sprintf("%s +\n`%s`", ret, line)
			line = ""
		}
	}

	return
}
