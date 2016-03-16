package engine

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"

	libuser "github.com/opencontainers/runc/libcontainer/user"
)

const (
	DefaultGroup = "crowley"
	// Errors message
	noGroupMatch = "no matching entries in group file"
)

var (
	ErrUnexpectedGroup = errors.New("group already exists")
)

func CreateGroup() (int, error) {

	gid, err := GroupID()

	if err != nil {
		return -1, err
	}

	_, err = libuser.LookupGid(gid)

	if err == nil {
		return -1, ErrUnexpectedGroup
	}

	if err == nil && err.Error() != noGroupMatch {
		return -1, err
	}

	if err = addgroup(gid); err != nil {
		return -1, err
	}

	return gid, nil
}

func addgroup(gid int) error {

	cmd := exec.Command("/usr/sbin/groupadd", "-g", strconv.FormatInt(int64(gid), 10), DefaultGroup)

	if buffer, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%s", string(buffer))
	}

	return nil
}
