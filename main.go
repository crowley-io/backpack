package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/crowley-io/backpack/engine"
	"github.com/jawher/mow.cli"
)

const (
	appName = "crowley-backpack"
	appDesc = "User management for crowley-pack build system."
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
			cli.Exit(Success)
		}

		conf, err := engine.ParseConfiguration(*path)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			cli.Exit(ErrParseConfiguration)
		}

		cli.Exit(handle(*conf))
	}

	backpack.Run(os.Args)
}
