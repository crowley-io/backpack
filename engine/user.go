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
)

var (
	// ErrUnexpectedUser is an error returned when the user already exists.
	ErrUnexpectedUser = errors.New("user already exists")
)

// NoUserLookup will check if the given uid or user name doesn't exists.
func NoUserLookup(uid int) error {

	// Get operating system-specific passwd reader-closer.
	passwd, err := libuser.GetPasswd()
	if err != nil {
		return err
	}
	defer passwd.Close()

	// Get the users.
	users, err := libuser.ParsePasswdFilter(passwd, func(u libuser.User) bool {
		return u.Uid == uid || u.Name == DefaultUser
	})

	// Check if an error has occurred.
	if err != nil {
		return err
	}

	// Check if no users entries found.
	if len(users) != 0 {
		return ErrUnexpectedUser
	}

	return nil
}

// CreateUser create a default user inside the container.
// The UID will be returned.
func CreateUser(gid int) (int, error) {

	uid, err := UserID()

	if err != nil {
		return -1, err
	}

	if err = NoUserLookup(uid); err != nil {
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
