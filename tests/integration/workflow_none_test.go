package integration_test

import (
	"bytes"
	"fmt"
	tasks "tasks/pkg/task_list"
	"tasks/pkg/test_utils"
	"testing"
)

func TestWorkflow_None(t *testing.T) {
	t.Run("go run .", func(t *testing.T) {
		task_list := []string{
			"One", "Two", "Another", "Task for Joe Doh",
		}
		actual := tasks.Print(task_list)
		expected := []string{
			"1. One",
			"2. Two",
			"3. Another",
			"4. Task for Joe Doh",
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

func TestWorkflow_Push(t *testing.T) {
	t.Run("push with empty list", func(t *testing.T) {
		expected := "[\"New task\"]"
		data := bytes.NewBuffer([]byte("[]"))
		var task_list = tasks.LoadState(data)
		task_list = tasks.Push(task_list, "New task")
		actual := tasks.SaveState(task_list)
		test_utils.AssertEqual(t, actual.String(), expected)
	})
	t.Run("push with items in list", func(t *testing.T) {
		expected := "[\"Task 1\",\"Task 2\",\"Task 3\"]"
		data := bytes.NewBuffer([]byte("[\"Task 1\",\"Task 2\"]"))
		var task_list = tasks.LoadState(data)
		task_list = tasks.Push(task_list, "Task 3")
		actual := tasks.SaveState(task_list)
		test_utils.AssertEqual(t, actual.String(), expected)
	})
	t.Run("push with list full", func(t *testing.T) {
		expected := "[\"Task 1\",\"Task 2\",\"Task 3\",\"Task 4\",\"Task 5\",\"Task 6\",\"Task 7\",\"Task 8\",\"Task 9\",\"Task 10\"]"
		data := bytes.NewBuffer([]byte("[\"Task 1\",\"Task 2\",\"Task 3\",\"Task 4\",\"Task 5\",\"Task 6\",\"Task 7\",\"Task 8\",\"Task 9\",\"Task 10\"]"))
		var task_list = tasks.LoadState(data)
		task_list = tasks.Push(task_list, "Another task")
		actual := tasks.SaveState(task_list)
		test_utils.AssertEqual(t, actual.String(), expected)
	})
}

func TestWorkflow_Pop(t *testing.T) {
	t.Run("pop with empty list", func(t *testing.T) {
		expected := "[]"
		data := bytes.NewBuffer([]byte("[]"))
		var task_list = tasks.LoadState(data)
		task_list, _ = tasks.Pop(task_list)
		actual := tasks.SaveState(task_list)
		test_utils.AssertEqual(t, actual.String(), expected)
	})
	t.Run("pop with one item in list", func(t *testing.T) {
		expected := "[]"
		data := bytes.NewBuffer([]byte("[\"One task\"]"))
		var task_list = tasks.LoadState(data)
		fmt.Println(task_list)
		task_list, _ = tasks.Pop(task_list)
		fmt.Println(task_list)
		actual := tasks.SaveState(task_list)
		fmt.Println(actual)
		test_utils.AssertEqual(t, actual.String(), expected)
	})
}
