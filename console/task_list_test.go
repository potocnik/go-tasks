package main

import (
	"tasks/test_utils"
	"testing"
)

func TestPrintTasks(t *testing.T) {
	t.Run("print tasks", func(t *testing.T) {
		actual := PrintTasks()
		expected := []string{
			"1. Task 1",
			"2. Task 2",
			"3. Task 3",
			"4. ",
			"5. ",
			"6. ",
			"7. ",
			"8. ",
			"9. ",
			"10. ",
		}
		test_utils.AssertEqualArray(t, actual, expected)
	})
}

func TestGetTasks(t *testing.T) {
	t.Run("get tasks", func(t *testing.T) {
		actual := GetTasks()
		expected := []string{
			"Task 1",
			"Task 2",
			"Task 3",
		}
		test_utils.AssertEqualArray(t, actual, expected)
	})
}
