package engine

import (
	"io"
)

func noLookup(read func() (io.ReadCloser, error), parse func(r io.Reader) error) error {

	var err error

	reader, err := read()
	if err != nil {
		return err
	}
	defer func() {
		if err2 := reader.Close(); err2 != nil && err == nil {
			err = err2
		}
	}()

	err = parse(reader)
	return err
}
