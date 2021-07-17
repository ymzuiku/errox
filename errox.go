package errox

import (
	"errors"
	"fmt"
	"runtime"
)

var Debug = false

func wrapStack(err error) error {
	pc, file, line, ok := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	if !ok {
		return errors.New("[bug]WrapError runtime.Caller(2) Fail")
	}
	return fmt.Errorf("%s:%d (Method %s)\n|- %w\n", file, line, f.Name(), err)
}

func Line() string {
	if Debug {
		pc, file, line, ok := runtime.Caller(1)
		f := runtime.FuncForPC(pc)
		if !ok {
			return "[bug]WrapError runtime.Caller(2) Fail"
		}
		return fmt.Sprintf("%s:%d (Method %s)\n", file, line, f.Name())
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