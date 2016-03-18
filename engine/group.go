package engine

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"

	libuser "github.com/opencontainers/runc/libcontainer/user"
)

const (
	// DefaultGroup define the group name
	DefaultGroup = "crowley"
)

var (
	// ErrUnexpectedGroup is an error returned when the group already exists.
	ErrUnexpectedGroup = errors.New("group already exists")
)

// NoGroupLookup will check if the given gid or group name doesn't exists.
func NoGroupLookup(gid int) error {

	// Get operating system-specific group reader-closer.
	group, err := libuser.GetGroup()
	if err != nil {
		return err
	}
	defer group.Close()

	// Get the groups.
	groups, err := libuser.ParseGroupFilter(group, func(g libuser.Group) bool {
		return g.Gid == gid || g.Name == DefaultGroup
	})

	// Check if an error has occurred.
	if err != nil {
		return err
	}

	// Check if no groups entries found.
	if len(groups) != 0 {
		return ErrUnexpectedGroup
	}

	return nil
}

// CreateGroup create a default group inside the container.
// The GID will be returned.
func CreateGroup() (int, error) {

	gid, err := GroupID()

	if err != nil {
		return -1, err
	}

	if err = NoGroupLookup(gid); err != nil {
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
