package exec

import (
	"os/exec"
	"syscall"
)

func Command(cmd string) (stdout string, exitcode int) {
	out, err := exec.Command("/bin/bash", "-c", cmd).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitcode = ws.ExitStatus()
		}
	}
	stdout = string(out[:])
	return
}
