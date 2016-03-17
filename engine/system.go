package engine

import (
	"os"
	"os/exec"
)

// CmdParser is a function which must return an executable Cmd from a given command.
type CmdParser func(string) *exec.Cmd

// Launch will execute all given commands, sequentially, and stop on the first command failure (if any).
func Launch(commands []string, create CmdParser) error {

	for _, command := range commands {

		cmd := create(command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
