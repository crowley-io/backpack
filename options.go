package main

import (
	"github.com/jawher/mow.cli"
)

func getBasicOption(app *cli.Cli) (version *bool) {

	// Force display of help in usage.
	app.Bool(cli.BoolOpt{
		Name:      "h help",
		Desc:      "Print usage and quit.",
		HideValue: true,
	})

	version = app.Bool(cli.BoolOpt{
		Name:      "v version",
		Desc:      "Print version information and quit.",
		HideValue: true,
	})

	return
}

func getConfigurationPathOption(app *cli.Cli) (path *string) {

	path = app.String(cli.StringOpt{
		Name:  "c configuration",
		Value: "packer.yml",
		Desc:  "Configuration file.",
	})

	return
}

func setRunCommandConfiguration(run *cli.Cmd) {

	run.Spec = "-- COMMAND [ARG...]"

	command := run.String(cli.StringArg{
		Name:      "COMMAND",
		Desc:      "Command to execute.",
		HideValue: true,
	})

	args := run.Strings(cli.StringsArg{
		Name:      "ARG",
		Desc:      "Command's arguments.",
		Value:     nil,
		HideValue: true,
	})

	run.Action = func() {
		cli.Exit(execute(*command, *args))
	}
}
