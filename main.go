package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/crowley-io/backpack/engine"
	"github.com/jawher/mow.cli"
)

// Application informations.
const (
	appName = "crowley-backpack"
	appDesc = "User management and command invoker for crowley-pack build system."
)

// Exit error codes
const (
	exitSuccess = iota
	exitErrParseConfiguration
	exitErrPreHookRuntime
	exitErrPostHookRuntime
	exitErrExecuteRuntime
	exitErrCreateGroup
	exitErrCreateUser
	exitErrUndefinedGroupEnv
	exitErrUndefinedUserEnv
	exitErrSetupUser
	exitErrLookPath
	exitErrSyscallExec
)

func init() {
	// Make sure we only have one process and that it runs on the main thread.
	runtime.GOMAXPROCS(1)
	runtime.LockOSThread()
}

func main() {

	backpack := cli.App(appName, appDesc)

	version := getBasicOption(backpack)
	path := getConfigurationPathOption(backpack)

	backpack.Command("run", "Run a command with another user and group ID.", setRunCommandConfiguration)

	backpack.Action = func() {

		if *version {
			fmt.Printf("%s %s\n", appName, engine.Version)
			cli.Exit(exitSuccess)
		}

		conf, err := engine.ParseConfiguration(*path)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			cli.Exit(exitErrParseConfiguration)
		}

		cli.Exit(handle(*conf))
	}

	if err := backpack.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
