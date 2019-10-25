package open

import (
	"os"
	"os/exec"
)

// URL opens a new browser window pointing to url.
func URL(url string) error {
	return openBrowser(url)
}

func runCmd(prog string, args ...string) error {
	cmd := exec.Command(prog, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	setFlags(cmd)
	return cmd.Run()
}
