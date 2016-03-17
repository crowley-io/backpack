package engine

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"

	libuser "github.com/opencontainers/runc/libcontainer/user"
)

const (
	// DefaultUser define the user name
	DefaultUser = "crowley"
	// Errors message
	noUserMatch = "no matching entries in passwd file"
)

var (
	// ErrUnexpectedUser is an error returned when the user already exists.
	ErrUnexpectedUser = errors.New("user already exists")
)

// CreateUser create a default user inside the container.
// The UID will be returned.
func CreateUser(gid int) (int, error) {

	uid, err := UserID()

	if err != nil {
		return -1, err
	}

	_, err = libuser.LookupUid(uid)

	if err == nil {
		return -1, ErrUnexpectedUser
	}

	if err == nil && err.Error() != noUserMatch {
		return -1, err
	}

	if err = adduser(uid, gid); err != nil {
		return -1, err
	}

	return uid, nil
}

func adduser(uid, gid int) error {

	cmd := exec.Command("/usr/sbin/useradd",
		"-g", strconv.FormatInt(int64(gid), 10),
		"-u", strconv.FormatInt(int64(uid), 10),
		"-m", DefaultUser,
	)

	if buffer, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%s", string(buffer))
	}

	return nil
}
