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
	ErrExecGidParamMatch = errors.New("gid doesn't match with given parameters")
	ErrExecUidParamMatch = errors.New("uid doesn't match with given parameters")
)

func Setup(uid, gid int) error {

	os.Unsetenv("HOME")
	id := fmt.Sprintf("%d:%d", uid, gid)

	// Set up defaults.
	defaultExecUser := user.ExecUser{
		Uid:  syscall.Getuid(),
		Gid:  syscall.Getgid(),
		Home: "/",
	}

	passwdPath, err := user.GetPasswdPath()
	if err != nil {
		return err
	}

	groupPath, err := user.GetGroupPath()
	if err != nil {
		return err
	}

	execUser, err := user.GetExecUserPath(id, &defaultExecUser, passwdPath, groupPath)
	if err != nil {
		return err
	}
	if execUser.Gid != gid {
		return ErrExecGidParamMatch
	}
	if execUser.Uid != uid {
		return ErrExecUidParamMatch
	}

	if err := syscall.Setgroups(execUser.Sgids); err != nil {
		return err
	}

	if err := system.Setgid(execUser.Gid); err != nil {
		return err
	}

	if err := system.Setuid(execUser.Uid); err != nil {
		return err
	}

	if err := os.Setenv("HOME", execUser.Home); err != nil {
		return err
	}

	return nil
}
