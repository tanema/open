package open

import "os/exec"

func open(input string) error {
	return exec.Command("xdg-open", input).Run()
}
