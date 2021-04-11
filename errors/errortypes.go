package errors

import (
	"github.com/pkg/errors"
)

type CanNotLoadList struct {
	error
}

func WrapCanNotLoadList(err error, format string, args ...interface{}) error {
	return &CanNotLoadList{errors.Wrapf(err, format, args...)}
}
