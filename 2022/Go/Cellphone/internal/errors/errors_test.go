package errors_test

import (
	"errors"
	"os"
	"testing"
)

func example3() *Errors {
	var err Errors
	err.AddError(errors.New("Example error 3"))
	return &err
}

func example2() *Errors {
	err := example3()
	err.AddError(errors.New("Example error 2"))
	return err
}

func example1() *Errors {
	err := example2()
	err.AddError(errors.New("Example error 1"))
	return err
}

func Test_String(t *testing.T) {
	os.WriteFile("out.txt", []byte(example1().Error()), 0644)
}
