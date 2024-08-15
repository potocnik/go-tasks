package unit_test

import (
	"errors"
	"tasks/pkg/utils"
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
		utils.Check(errorUnderTest)
	})
	t.Run("no error should result in no panic", func(t *testing.T) {
		utils.Check(nil)
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
		utils.CheckWithMessage(errorUnderTest, "Example")
	})
	t.Run("no error should result in no panic", func(t *testing.T) {
		utils.CheckWithMessage(nil, "Example")
	})
}
