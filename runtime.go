package main

import (
	"fmt"
	"github.com/crowley-io/backpack/engine"
	"os"
	"os/exec"
	"syscall"
	//"strings"
)

func handle(c engine.Configuration) int {

	// Create user and group.
	gid, err := engine.CreateGroup()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ErrCreateGroup
	}

	if _, err = engine.CreateUser(gid); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ErrCreateUser
	}

	// Execute pre-hooks.
	if err := engine.Launch(c.Prehooks(), createHookCmd); err != nil {
		return ErrPreHookRuntime
	}

	// Execute every command as required user.
	if err := engine.Launch(c.Execute(), createRunCmd); err != nil {
		fmt.Fprintln(os.Stderr, "Cannot execute: ", err)
		return ErrExecuteRuntime
	}

	// Execute post-hooks.
	if err := engine.Launch(c.Posthooks(), createHookCmd); err != nil {
		return ErrPostHookRuntime
	}

	return Success
}

func execute(command string, args []string) int {

	gid, err := engine.GroupID()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ErrUndefinedGroupEnv
	}

	uid, err := engine.UserID()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ErrUndefinedUserEnv
	}

	if err := engine.Setup(uid, gid); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ErrSetupUser
	}

	path, err := exec.LookPath(command)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ErrLookPath
	}

	args = append([]string{command}, args...)

	if err = syscall.Exec(path, args, os.Environ()); err != nil {
		fmt.Fprintln(os.Stderr, "Cannot execute: ", err)
		return ErrSyscallExec
	}

	return 0
}

func createHookCmd(cmd string) *exec.Cmd {
	return exec.Command("bash", "-c", cmd)
}

func createRunCmd(cmd string) *exec.Cmd {
	return exec.Command(os.Args[0], "run", "bash", "-c", cmd)
}
