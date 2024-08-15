package tasks

import (
	tasks "tasks/pkg/task_list"
	"tasks/pkg/test_utils"
	"testing"
)

func TestPrintTasks(t *testing.T) {
	t.Run("print tasks", func(t *testing.T) {
		actual := tasks.PrintTasks()
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
		actual := tasks.GetTasks()
		expected := []string{
			"Task 1",
			"Task 2",
			"Task 3",
		}
		test_utils.AssertEqualArray(t, actual, expected)
	})
}

func TestWriteTasks(t *testing.T) {
	t.Run("write tasks", func(t *testing.T) {
		actual := tasks.WriteTasks()
		expected := "[\"Task 1\",\"Task 2\",\"Task 3\"]"
		test_utils.AssertEqual(t, actual.String(), expected)
	})
}
