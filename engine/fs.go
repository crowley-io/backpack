package engine

import (
	"os"
)

// ChangeDirectory will change the current working directory to the required one.
// Also, it will update the "PWD" environment variable.
func ChangeDirectory() error {
	return changeDirectory(WorkingDirectory, os.Chdir, os.Setenv)
}

func changeDirectory(directory func() (path string, err error),
	chdir func(path string) error,
	setenv func(key, value string) error) error {

	dir, err := directory()
	if err != nil {
		return err
	}

	if err = chdir(dir); err != nil {
		return err
	}

	if err = setenv("PWD", dir); err != nil {
		return err
	}

	return nil
}
