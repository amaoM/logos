package error

import (
	"errors"
	"fmt"
)

func logos1() error {
	err := doErrorUsingErrors()
	if err != nil {
		return err
	}
	return nil
}

func doErrorUsingErrors() error {
	return errors.New("Errors")
}

func logos2() error {
	err := doErrorUsingFmtError()
	if err != nil {
		return err
	}
	return nil
}

func doErrorUsingFmtError() error {
	return fmt.Errorf("%s", "FmtError")
}

func logos3() error {
	err := doErrorUsingErrorStruct()
	if err != nil {
		switch e := err.(type) {
		case *MyError:
			return e
		default:
			return errors.New("Default error")
		}
	}
	return nil
}

// MyError ...
type MyError struct {
	Message string
}

// error interface have Error method.
func (me *MyError) Error() string {
	return fmt.Sprintf("%s", me.Message)
}

func doErrorUsingErrorStruct() error {
	return &MyError{Message: "This step is doErrorUsingErrorStruct"}
}
