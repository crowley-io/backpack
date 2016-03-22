package engine

import (
	"os"
)

// ChangeDirectory will change the current working directory to the required one.
// Also, it will update the "PWD" environment variable.
func ChangeDirectory() error {

	dir, err := WorkingDirectory()
	if err != nil {
		return err
	}

	if err = os.Chdir(dir); err != nil {
		return err
	}

	if err = os.Setenv("PWD", dir); err != nil {
		return err
	}

	return nil
}
