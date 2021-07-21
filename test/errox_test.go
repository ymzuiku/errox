package test

import (
	"regexp"
	"testing"

	"github.com/ymzuiku/errox"
	"github.com/ymzuiku/errox/test/othertest"
	"github.com/ymzuiku/so"
)

func errorDog() error {
	err := errox.New("dog")
	return err
}

func TestShowLine(t *testing.T) {
	errox.Debug = true
	t.Run("test in this file", func(t *testing.T) {
		err := errorDog()
		so.True(t, regexp.MustCompile(`test/errox_test.go`).Match([]byte(err.Error())))
	})

	t.Run("test in other file", func(t *testing.T) {
		err := othertest.OtherError()
		so.True(t, regexp.MustCompile(`errox/test/othertest/othertest.go:12`).Match([]byte(err.Error())))
		so.True(t, regexp.MustCompile(`errox/test/othertest/othertest.go:6`).Match([]byte(err.Error())))
	})
}
