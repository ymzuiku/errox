package errox

import (
	"errors"
	"fmt"
	"runtime"
)

var Debug = false

func wrapStack(err error) error {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return errors.New("[bug]WrapError runtime.Caller(2) Fail")
	}
	return fmt.Errorf("%s:%d\n%w\n", file, line, err)
}

func Line() string {
	if Debug {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			return "[bug]WrapError runtime.Caller(2) Fail"
		}
		return fmt.Sprintf("%s:%d\n", file, line)
	}
	return ""
}

func Stack(err error) error {
	return wrapStack(err)
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	if Debug {
		return wrapStack(err)
	}
	return err
}

func New(msg string) error {
	if Debug {
		return wrapStack(errors.New(msg))
	}
	return errors.New(msg)
}

func Errorf(format string, a ...interface{}) error {
	if Debug {
		return wrapStack(fmt.Errorf(format, a...))
	}
	return fmt.Errorf(format, a...)
}
