package engine

import (
	"errors"
	"fmt"
	"os"
	"syscall"

	"github.com/opencontainers/runc/libcontainer/system"
	"github.com/opencontainers/runc/libcontainer/user"
)

var (
	// ErrExecGIDParamMatch is an error returned when the user's gid is not the required/requested one.
	ErrExecGIDParamMatch = errors.New("gid doesn't match with given parameters")
	// ErrExecUIDParamMatch is an error returned when the user's uid is not the required/requested one.
	ErrExecUIDParamMatch = errors.New("uid doesn't match with given parameters")
)

func getUserPaths() (passwdPath, groupPath string, err error) {

	if passwdPath, err = user.GetPasswdPath(); err != nil {
		return "", "", err
	}

	if groupPath, err = user.GetGroupPath(); err != nil {
		return "", "", err
	}

	return
}

func getUserExec(uid, gid int) (*user.ExecUser, error) {

	id := fmt.Sprintf("%s:%s", DefaultUser, DefaultGroup)

	// Set up defaults.
	defaultExecUser := user.ExecUser{
		Uid:  syscall.Getuid(),
		Gid:  syscall.Getgid(),
		Home: "/",
	}

	passwdPath, groupPath, err := getUserPaths()
	if err != nil {
		return nil, err
	}

	execUser, err := user.GetExecUserPath(id, &defaultExecUser, passwdPath, groupPath)
	if err != nil {
		return nil, err
	}

	if execUser.Gid != gid {
		return nil, ErrExecGIDParamMatch
	}
	if execUser.Uid != uid {
		return nil, ErrExecUIDParamMatch
	}

	return execUser, nil
}

func setUserExecContext(u *user.ExecUser) error {

	if err := syscall.Setgroups(u.Sgids); err != nil {
		return err
	}

	if err := system.Setgid(u.Gid); err != nil {
		return err
	}

	if err := system.Setuid(u.Uid); err != nil {
		return err
	}

	if err := os.Setenv("HOME", u.Home); err != nil {
		return err
	}

	return nil
}

// Setup will change running process to given uid and gid.
// NOTE: You can't rollback to the previous user if its successful.
func Setup(uid, gid int) error {

	// Clear HOME since it will set later...
	if err := os.Unsetenv("HOME"); err != nil {
		return err
	}

	execUser, err := getUserExec(uid, gid)
	if err != nil {
		return err
	}

	if err := setUserExecContext(execUser); err != nil {
		return err
	}

	return nil
}
