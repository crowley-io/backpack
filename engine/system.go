package engine

import (
	"os"
	"os/exec"
	"syscall"
)

// Launch will execute the given command without forking.
func Launch(args []string) error {
	return launch(args, exec.LookPath, syscall.Exec)
}

func launch(args []string,
	lookpath func(cmd string) (path string, err error),
	exec func(arg0 string, argv []string, env []string) error) error {

	// Find command path.
	path, err := lookpath(args[0])
	if err != nil {
		return err
	}

	// Execute given command.
	return exec(path, args, os.Environ())
}
