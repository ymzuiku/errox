package errox

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
)

func wrapStack(err error) error {
	_, file, line, ok := runtime.Caller(2)

	if !ok {
		return errors.New("[errox]WrapError runtime.Caller(2) Fail")
	}

	return fmt.Errorf("\n%s:%d %w", file, line, err)
}

func lineStack() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "[bug]WrapError runtime.Caller(2) Fail"
	}
	return fmt.Sprintf("%s:%d", file, line)
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
		return err
	}
	if Debug {
		return wrapStack(err)
	}
	return err
}

// Same errox.Wrap, add format string
func Wrapf(err error, format string, a ...interface{}) error {
	if err == nil {
		return fmt.Errorf("%v %w", fmt.Sprintf(format, a...), err)
	}
	if Debug {
		return wrapStack(fmt.Errorf("%v %w", fmt.Sprintf(format, a...), err))
	}
	return fmt.Errorf("%v %w", fmt.Sprintf(format, a...), err)
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
		return lineStack()
	}
	return ""
}

// log.Fatalf and add stack line
func Fatalf(format string, v ...interface{}) {
	if Debug {
		fmt.Printf(lineStack()+" "+format, v...)
		os.Exit(1)
		return
	}
	log.Fatalf(format, v...)
}

// log.Panicf and add stack line
func Panicf(format string, v ...interface{}) {
	if Debug {
		panic(fmt.Sprintf(lineStack()+" "+format, v...))
	}
	log.Panicf(format, v...)
}

// log.Printf and add stack line
func Printf(format string, v ...interface{}) {
	if Debug {
		fmt.Printf(lineStack()+" "+format, v...)
		return
	}
	log.Printf(format, v...)
}
