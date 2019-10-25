package open

import "os/exec"

func openBrowser(url string) error {
	return runCmd("xdg-open", url)
}

func setFlags(cmd *exec.Cmd) {}
