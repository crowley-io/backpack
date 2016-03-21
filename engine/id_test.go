package engine

import (
	"fmt"
	"testing"
)

func TestGroupID(t *testing.T) {
	checkIDMatch(t, GroupEnv, 1005, GroupID)
}

func TestUserID(t *testing.T) {
	checkIDMatch(t, UserEnv, 1002, UserID)
}

func TestWorkingDirectory(t *testing.T) {

	e := "/home/user"

	setEnv(t, DirectoryEnv, "/home/user")
	p, err := WorkingDirectory()

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if p != e {
		t.Fatalf("expected path %s but received %s", e, p)
	}

	unsetEnv(t, DirectoryEnv)

}

func TestGroupIDWithSyntaxError(t *testing.T) {
	setEnv(t, GroupEnv, "foo")
	defer unsetEnv(t, GroupEnv)
	checkIDError(t, ErrGroupEnvSyntaxError, GroupID)
}

func TestUserIDWithSyntaxError(t *testing.T) {
	setEnv(t, UserEnv, "foo")
	defer unsetEnv(t, UserEnv)
	checkIDError(t, ErrUserEnvSyntaxError, UserID)
}

func TestGroupIDWithUndefinedEnv(t *testing.T) {
	checkIDError(t, ErrUndefinedGroupEnv, GroupID)
}

func TestUserIDWithUndefinedEnv(t *testing.T) {
	checkIDError(t, ErrUndefinedUserEnv, UserID)
}

func TestWorkingDirectoryWithUndefinedEnv(t *testing.T) {

	e := ErrUndefinedDirectoryEnv
	_, err := WorkingDirectory()

	if err == nil {
		t.Fatalf("an error was expected")
	}

	if err != e {
		t.Fatalf("expected error '%s' but got '%s'", e, err)
	}

}

func setEnv(t *testing.T, key, value string) {
	if err := os.Setenv(key, value); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func unsetEnv(t *testing.T, key string) {
	if err := os.Unsetenv(key); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func checkIDMatch(t *testing.T, key string, e int, callback func() (int, error)) {

	setEnv(t, key, fmt.Sprint(e))

	i, err := callback()

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if i != e {
		t.Fatalf("expected id %d but received %d", e, i)
	}

	unsetEnv(t, key)
}

func checkIDError(t *testing.T, e error, callback func() (int, error)) {

	_, err := callback()

	if err == nil {
		t.Fatalf("an error was expected")
	}

	if err != e {
		t.Fatalf("expected error '%s' but got '%s'", e, err)
	}
}
