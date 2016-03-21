package engine

import (
	"errors"
	"fmt"
	"testing"
)

func TestChangeDirectory(t *testing.T) {

	directoryCalled := false
	chdirCalled := false
	setenvCalled := false
	expectedKey := "PWD"
	expectedPath := "/usr/local/app"

	directory := func() (string, error) {
		directoryCalled = true
		return expectedPath, nil
	}

	chdir := func(path string) error {
		chdirCalled = true
		if path != expectedPath {
			return fmt.Errorf("unexpected path: %s", path)
		}
		return nil
	}

	setenv := func(key, value string) error {
		setenvCalled = true
		if key != expectedKey {
			return fmt.Errorf("unexpected env's key: %s", key)
		}
		if value != expectedPath {
			return fmt.Errorf("unexpected env's value: %s", value)
		}
		return nil
	}

	err := changeDirectory(directory, chdir, setenv)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if directoryCalled == false {
		t.Fatalf("call of directory lambda was expected")
	}

	if chdirCalled == false {
		t.Fatalf("call of chdir lambda was expected")
	}

	if setenvCalled == false {
		t.Fatalf("call of setenv lambda was expected")
	}
}

func TestChangeDirectoryWithPathError(t *testing.T) {

	e := errors.New("unknown path")

	directory := func() (string, error) {
		return "", e
	}

	chdir := func(path string) error {
		err := errors.New("chdir call was unexpected")
		t.Fatal(err)
		return err
	}

	setenv := func(key, value string) error {
		err := errors.New("setenv call was unexpected")
		t.Fatal(err)
		return err
	}

	err := changeDirectory(directory, chdir, setenv)

	if err == nil {
		t.Fatal("an error was expected")
	}

	if err != e {
		t.Fatalf("expected error '%s' but received '%s'", e, err)
	}

}

func TestChangeDirectoryWithChdirError(t *testing.T) {

	directoryCalled := false
	chdirCalled := false
	e := errors.New("chdir: permission denied")

	directory := func() (string, error) {
		directoryCalled = true
		return "/usr/local/app", nil
	}

	chdir := func(path string) error {
		chdirCalled = true
		return e
	}

	setenv := func(key, value string) error {
		err := errors.New("setenv call was unexpected")
		t.Fatal(err)
		return err
	}

	err := changeDirectory(directory, chdir, setenv)

	if err == nil {
		t.Fatal("an error was expected")
	}

	if err != e {
		t.Fatalf("expected error '%s' but received '%s'", e, err)
	}

	if directoryCalled == false {
		t.Fatalf("call of directory lambda was expected")
	}

	if chdirCalled == false {
		t.Fatalf("call of chdir lambda was expected")
	}

}

func TestChangeDirectoryWithSetenvError(t *testing.T) {

	directoryCalled := false
	chdirCalled := false
	setenvCalled := false
	e := errors.New("export: permission denied")

	directory := func() (string, error) {
		directoryCalled = true
		return "/usr/local/app", nil
	}

	chdir := func(path string) error {
		chdirCalled = true
		return nil
	}

	setenv := func(key, value string) error {
		setenvCalled = true
		return e
	}

	err := changeDirectory(directory, chdir, setenv)

	if err == nil {
		t.Fatal("an error was expected")
	}

	if err != e {
		t.Fatalf("expected error '%s' but received '%s'", e, err)
	}

	if directoryCalled == false {
		t.Fatalf("call of directory lambda was expected")
	}

	if chdirCalled == false {
		t.Fatalf("call of chdir lambda was expected")
	}

	if setenvCalled == false {
		t.Fatalf("call of setenv lambda was expected")
	}

}
