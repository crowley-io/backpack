package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/crowley-io/backpack/engine"
)

// Exit error codes
const (
	exitSuccess = iota
	exitErrCreateGroup
	exitErrCreateUser
	exitErrWorkingDirectory
	exitErrSetupUser
	exitErrSyscallExec
)

func init() {
	// Make sure we only have one process and that it runs on the main thread.
	runtime.GOMAXPROCS(1)
	runtime.LockOSThread()
}

func main() {

	backpack := newApp(
		"crowley-backpack", "User management and command invoker for crowley-pack build system.", engine.Version,
	)

	backpack.Handler = handle
	backpack.Run(os.Args)

}

func handle(args []string) int {

	// Create user and group.
	gid, err := engine.CreateGroup()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErrCreateGroup
	}

	uid, err := engine.CreateUser(gid)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErrCreateUser
	}

	// Change working directory
	err = engine.ChangeDirectory()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErrWorkingDirectory
	}

	// Switch to user.
	if err = engine.Setup(uid, gid); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErrSetupUser
	}

	// Execute given command.
	if err = engine.Launch(args); err != nil {
		fmt.Fprintln(os.Stderr, "Execution error: ", err)
		return exitErrSyscallExec
	}

	return exitSuccess
}
