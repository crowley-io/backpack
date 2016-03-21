package engine

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestLaunch(t *testing.T) {

	args := []string{"ls", "-l"}

	lookpath := func(cmd string) (path string, err error) {
		if cmd == "ls" {
			return "/usr/bin/ls", nil
		}
		return "", fmt.Errorf("unexpected command: %s", cmd)
	}

	exec := func(arg0 string, argv []string, env []string) error {
		if arg0 == "/usr/bin/ls" && reflect.DeepEqual(argv, args) && len(env) > 0 {
			t.Logf("run: %s", strings.Join(argv, " "))
			return nil
		}
		return fmt.Errorf("unexpected configuration: %s %+v %+v", arg0, argv, env)
	}

	err := launch(args, lookpath, exec)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

}

func TestLaunchWithLookpathError(t *testing.T) {

	args := []string{"ls", "-l"}
	e := errors.New("no ls found in PATH")

	lookpath := func(cmd string) (path string, err error) {
		return "", e
	}

	exec := func(arg0 string, argv []string, env []string) error {
		err := errors.New("exec call was unexpected")
		t.Fatal(err)
		return err
	}

	err := launch(args, lookpath, exec)

	if err == nil {
		t.Fatal("an error was expected error")
	}

	if err != e {
		t.Fatalf("expected error '%s' but received '%s'", e, err)
	}

}

func TestLaunchWithExecError(t *testing.T) {

	args := []string{"mkdir", "/dev/foo"}
	e := errors.New("mkdir: permission denied")

	lookpath := func(cmd string) (path string, err error) {
		if cmd == "mkdir" {
			return "/usr/bin/mkdir", nil
		}
		return "", fmt.Errorf("unexpected command: %s", cmd)
	}

	exec := func(arg0 string, argv []string, env []string) error {
		return e
	}

	err := launch(args, lookpath, exec)

	if err == nil {
		t.Fatal("an error was expected error")
	}

	if err != e {
		t.Fatalf("expected error '%s' but received '%s'", e, err)
	}

}
