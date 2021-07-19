package errox

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/fatih/color"
)

func wrapStack(err error) error {
	_, file, line, ok := runtime.Caller(2)

	if !ok {
		return errors.New(color.RedString("[errox]WrapError runtime.Caller(2) Fail"))
	}

	return fmt.Errorf("\n%s%s%s%w", color.HiRedString(file), color.HiRedString(":%d", line), color.HiRedString(" : "), err)
}

// Change debug bool, base is false
var Debug = false

// Create a error and add stack info string when errox.Debug = true
func New(msg string) error {
	if Debug {
		return wrapStack(errors.New(msg))
	}
	return errors.New(msg)
}

// Add stack info to error when errox.Debug = true
func Wrap(err error) error {
	if err == nil {
		return nil
	}
	if Debug {
		return wrapStack(err)
	}
	return err
}

// Same errox.Wrap, add format string
func Wrapf(err error, format string, a ...interface{}) error {
	if err == nil {
		return nil
	}
	if Debug {
		return wrapStack(fmt.Errorf("%v %w", fmt.Sprintf(format, a...), err))
	}
	return err
}

// Add stack info and fmt.Errorf when errox.Debug = true
func Errorf(format string, a ...interface{}) error {
	if Debug {
		return wrapStack(fmt.Errorf(format, a...))
	}
	return fmt.Errorf(format, a...)
}

// Make stack info string when errox.Debug = true
func Line() string {
	if Debug {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			return "[bug]WrapError runtime.Caller(2) Fail"
		}
		return fmt.Sprintf("\n%s:%d", file, line)
	}
	return ""
}
