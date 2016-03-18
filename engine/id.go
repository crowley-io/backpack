package engine

import (
	"fmt"
	"os"
	"strconv"
)

const (
	// UserEnv is the environment variable name given by pack to define the user id.
	UserEnv = "CROWLEY_PACK_USER"
	// GroupEnv is the environment variable name given by pack to define the group id.
	GroupEnv = "CROWLEY_PACK_GROUP"
	// DirectoryEnv is the environment variable name given by pack to define the working
	// directory inside the container.
	DirectoryEnv = "CROWLEY_PACK_DIRECTORY"
)

var (
	// ErrUndefinedUserEnv is an error returned when the user environment variable is undefined.
	ErrUndefinedUserEnv = fmt.Errorf("'%s' environment variable is undefined", UserEnv)
	// ErrUndefinedGroupEnv is an error returned when the group environment variable is undefined.
	ErrUndefinedGroupEnv = fmt.Errorf("'%s' environment variable is undefined", GroupEnv)
	// ErrUndefinedDirectoryEnv is an error returned when the directory environment variable is undefined.
	ErrUndefinedDirectoryEnv = fmt.Errorf("'%s' environment variable is undefined", DirectoryEnv)
)

// GroupID returns the required group's id for this process.
func GroupID() (int, error) {
	return parseID(GroupEnv, ErrUndefinedGroupEnv)
}

// UserID returns the required user's id for this process.
func UserID() (int, error) {
	return parseID(UserEnv, ErrUndefinedUserEnv)
}

// WorkingDirectory returns the required working directory for this process.
func WorkingDirectory() (string, error) {

	s := os.Getenv(DirectoryEnv)

	if s == "" {
		return "", ErrUndefinedDirectoryEnv
	}

	return s, nil
}

func parseID(env string, err error) (int, error) {

	s := os.Getenv(env)

	if s == "" {
		return 0, err
	}

	id, err := parseInt(s)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func parseInt(s string) (int, error) {

	i, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		return 0, err
	}

	return int(i), nil
}
