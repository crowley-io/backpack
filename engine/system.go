package engine

import (
	"os"
	"os/exec"
	"syscall"
)

// Launch will execute the given command without forking.
func Launch(args []string) error {

	// Find command path.
	path, err := exec.LookPath(args[0])
	if err != nil {
		return err
	}

	// Execute given command.
	return syscall.Exec(path, args, os.Environ())
}
