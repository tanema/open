package open

import "os/exec"

func open(input string) error {
	return exec.Command("open", input).Run()
}
