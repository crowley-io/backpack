package engine

import (
	"os"
	"os/exec"
)

type CmdParser func(string) *exec.Cmd

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
