package othertest

import "github.com/ymzuiku/errox"

func sub() error {
	err := errox.New("Sub error")
	return err
}

func OtherError() error {
	subError := sub()
	err := errox.Wrap(subError)

	return err
}
