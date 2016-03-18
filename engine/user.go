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
	// DefaultUser define the user name
	DefaultUser = "crowley"
)

var (
	// ErrUnexpectedUser is an error returned when the user already exists.
	ErrUnexpectedUser = errors.New("user already exists")
)

func noUserLookup(uid int) error {

	reader := func() (io.ReadCloser, error) {
		return libuser.GetPasswd()
	}

	parser := func(r io.Reader) error {

		users, err := libuser.ParsePasswdFilter(r, func(u libuser.User) bool {
			return u.Uid == uid || u.Name == DefaultUser
		})

		if err == nil && len(users) != 0 {
			return ErrUnexpectedGroup
		}

		return nil
	}

	return noLookup(reader, parser)
}

// CreateUser create a default user inside the container.
// The UID will be returned.
func CreateUser(gid int) (int, error) {

	uid, err := UserID()

	if err != nil {
		return -1, err
	}

	if err = noUserLookup(uid); err != nil {
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
