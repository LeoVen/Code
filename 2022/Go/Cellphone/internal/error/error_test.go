package error

import (
	"errors"
	"os"
	"testing"
)

func example3() *Error {
	var err Error
	err.AddError(errors.New("Example error 3"))
	return &err
}

func example2() *Error {
	err := example3()
	err.AddError(errors.New("Example error 2"))
	return err
}

func example1() *Error {
	err := example2()
	err.AddError(errors.New("Example error 1"))
	return err
}

func Test_String(t *testing.T) {
	os.WriteFile("out.txt", []byte(example1().Error()), 0644)
}
