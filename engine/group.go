package engine

import (
	"errors"
	"fmt"
	"io"
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

func noGroupLookup(gid int) error {

	reader := func() (io.ReadCloser, error) {
		return libuser.GetGroup()
	}

	parser := func(r io.Reader) error {

		groups, err := libuser.ParseGroupFilter(r, func(g libuser.Group) bool {
			return g.Gid == gid || g.Name == DefaultGroup
		})

		if err == nil && len(groups) != 0 {
			return ErrUnexpectedGroup
		}

		return nil
	}

	return noLookup(reader, parser)
}

// CreateGroup create a default group inside the container.
// The GID will be returned.
func CreateGroup() (int, error) {

	gid, err := GroupID()

	if err != nil {
		return -1, err
	}

	if err = noGroupLookup(gid); err != nil {
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
