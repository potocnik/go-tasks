package unit_test

import (
	"errors"
	error "tasks/pkg/utils/errors"
	"testing"
)

func TestCheck(t *testing.T) {
	t.Run("error should result in panic", func(t *testing.T) {
		errorUnderTest := errors.New("Test error")
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		error.Check(errorUnderTest)
	})
	t.Run("no error should result in no panic", func(t *testing.T) {
		error.Check(nil)
	})
}

func TestCheckWithMessage(t *testing.T) {
	t.Run("error should result in panic", func(t *testing.T) {
		errorUnderTest := errors.New("Test error")
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		error.CheckWithMessage(errorUnderTest, "Example")
	})
	t.Run("no error should result in no panic", func(t *testing.T) {
		error.CheckWithMessage(nil, "Example")
	})
}
