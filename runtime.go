package main

import (
	"fmt"
	"github.com/crowley-io/backpack/engine"
	"os"
	"os/exec"
	"syscall"
)

// handle is the backpack's supervisor:
//   - First, it will create required user and group (defined by crowley-pack's environment variable).
//   - Then, if any pre-hooks are defined, it will execute them in a child process.
//   - After that, it will launch itself - also in a child process - in order to execute every command defined in
//     the configuration file. Each command will be executed as the user created in the first step...
//   - Finally, if any post-hooks are defined, it will execute them in a child process.
func handle(path string) int {

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

	// Change working directory
	dir, err := engine.WorkingDirectory()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErrUndefinedDirectoryEnv
	}

	if err = os.Chdir(dir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErrWorkingDirectory
	}

	// Parse configuration
	c, err := engine.ParseConfiguration(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErrParseConfiguration
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

// execute is the backpack's command wrapper:
// It will switch to the required user (defined by crowley-pack's environment variable) and then execute the
// specified command. Also, it will be no longer resident or involved in the process lifecycle at all once the
// command is started.
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

	// Switch to user.
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

	// Execute given command.
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
