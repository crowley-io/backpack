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
		return exitErrCreateGroup
	}

	if _, err = engine.CreateUser(gid); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErrCreateUser
	}

	// Execute pre-hooks.
	if err := engine.Launch(c.Prehooks(), createHookCmd); err != nil {
		return exitErrPreHookRuntime
	}

	// Execute every command as required user.
	if err := engine.Launch(c.Execute(), createRunCmd); err != nil {
		fmt.Fprintln(os.Stderr, "Cannot execute: ", err)
		return exitErrExecuteRuntime
	}

	// Execute post-hooks.
	if err := engine.Launch(c.Posthooks(), createHookCmd); err != nil {
		return exitErrPostHookRuntime
	}

	return exitSuccess
}

func execute(command string, args []string) int {

	gid, err := engine.GroupID()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErrUndefinedGroupEnv
	}

	uid, err := engine.UserID()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErrUndefinedUserEnv
	}

	if err = engine.Setup(uid, gid); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErrSetupUser
	}

	path, err := exec.LookPath(command)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErrLookPath
	}

	args = append([]string{command}, args...)

	if err = syscall.Exec(path, args, os.Environ()); err != nil {
		fmt.Fprintln(os.Stderr, "Cannot execute: ", err)
		return exitErrSyscallExec
	}

	return exitSuccess
}

func createHookCmd(cmd string) *exec.Cmd {
	return exec.Command("bash", "-c", cmd)
}

func createRunCmd(cmd string) *exec.Cmd {
	return exec.Command(os.Args[0], "run", "bash", "-c", cmd)
}
