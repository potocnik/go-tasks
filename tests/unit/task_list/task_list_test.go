package unit_test

import (
	"bytes"
	tasks "tasks/pkg/task_list"
	"tasks/pkg/test_utils"
	"testing"
)

func TestPrint(t *testing.T) {
	t.Run("print tasks", func(t *testing.T) {
		actual := tasks.Print([]string{"Task 1", "Task 2", "Task 3"})
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

func TestPush(t *testing.T) {
	t.Run("push a task", func(t *testing.T) {
		actual := tasks.Push([]string{"Task 1", "Task 2", "Task 3"}, "New task")
		expected := []string{"Task 1", "Task 2", "Task 3", "New task"}
		test_utils.AssertEqualArray(t, actual, expected)
	})
}

func TestPop(t *testing.T) {
	t.Run("pop a task", func(t *testing.T) {
		actual, _ := tasks.Pop([]string{"Task 1", "Task 2", "Task 3"})
		expected := []string{"Task 2", "Task 3"}
		test_utils.AssertEqualArray(t, actual, expected)
	})
}

func TestLoadState(t *testing.T) {
	t.Run("save state", func(t *testing.T) {
		data := bytes.NewBuffer([]byte("[\"Task 1\",\"Task 2\",\"Task 3\"]"))
		actual := tasks.LoadState(data)
		expected := []string{"Task 1", "Task 2", "Task 3"}
		test_utils.AssertEqualArray(t, actual, expected)
	})
}

func TestSaveState(t *testing.T) {
	t.Run("load state", func(t *testing.T) {
		actual := tasks.SaveState([]string{"Task 1", "Task 2", "Task 3"})
		expected := "[\"Task 1\",\"Task 2\",\"Task 3\"]"
		test_utils.AssertEqual(t, actual.String(), expected)
	})
}
