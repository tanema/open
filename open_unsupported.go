// +build !linux,!windows,!darwin,!openbsd

package open

import (
	"fmt"
	"os/exec"
	"runtime"
)

func openBrowser(url string) error {
	return fmt.Errorf("openBrowser: unsupported operating system: %v", runtime.GOOS)
}

func setFlags(cmd *exec.Cmd) {}
