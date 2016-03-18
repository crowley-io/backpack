package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type app struct {
	name        string
	description string
	version     string
	Handler     func(args []string) int
}

func newApp(name, description, version string) *app {
	return &app{
		name:        name,
		description: description,
		version:     version,
		Handler:     func(args []string) int { return exitSuccess },
	}
}

func (e *app) Version(w io.Writer) {
	fmt.Fprintf(w, "%s %s\n", e.name, e.version)
}

func (e *app) Usage(w io.Writer) {

	fmt.Fprintln(w)
	fmt.Fprintf(w, "Usage: %s command [args]\n", e.name)
	fmt.Fprintf(w, "       %s --version\n", e.name)
	fmt.Fprintln(w)
	fmt.Fprintf(w, "%s\n", e.description)
	fmt.Fprintln(w)
	fmt.Fprintf(w, "Arguments:\n")
	fmt.Fprintf(w, "  command      Command to execute\n")
	fmt.Fprintf(w, "  args         Command's arguments\n")
	fmt.Fprintln(w)
	fmt.Fprintf(w, "Options:\n")
	fmt.Fprintf(w, "  -h, --help       Print usage and quits\n")
	fmt.Fprintf(w, "  -v, --version    Print version information and quits\n")

}

func (e *app) Run(args []string) {

	if len(args) < 2 {
		e.Usage(os.Stderr)
		os.Exit(255)
	}

	l := strings.TrimSpace(strings.Join(args[1:], " "))

	switch l {
	case "-v", "--version":
		e.Version(os.Stdout)
	case "-h", "--help":
		e.Usage(os.Stdout)
	default:
		os.Exit(e.Handler(args[1:]))
	}
}
