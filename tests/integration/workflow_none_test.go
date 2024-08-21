package integration_test

import (
	"bytes"
	tasks "tasks/pkg/task_list"
	"tasks/pkg/test_utils"
	"testing"
)

const FULL_TASKS = "[\"Task 1\",\"Task 2\",\"Task 3\",\"Task 4\",\"Task 5\",\"Task 6\",\"Task 7\",\"Task 8\",\"Task 9\",\"Task 10\"]"
const EMPTY_TASKS = "[]"
const ONE_TASK = "[\"One Task\"]"

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
		expected := FULL_TASKS
		data := bytes.NewBuffer([]byte(FULL_TASKS))
		var task_list = tasks.LoadState(data)
		task_list = tasks.Push(task_list, "Another task")
		actual := tasks.SaveState(task_list)
		test_utils.AssertEqual(t, actual.String(), expected)
	})
	t.Run("push multiple items", func(t *testing.T) {
		expected := FULL_TASKS
		data := bytes.NewBuffer([]byte("[\"Task 1\",\"Task 2\",\"Task 3\",\"Task 4\",\"Task 5\",\"Task 6\",\"Task 7\",\"Task 8\"]"))
		var task_list = tasks.LoadState(data)
		task_list = tasks.Push(task_list, "Task 9")
		task_list = tasks.Push(task_list, "Task 10")
		task_list = tasks.Push(task_list, "Task 11")
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
		task_list, _ = tasks.Pop(task_list)
		actual := tasks.SaveState(task_list)
		test_utils.AssertEqual(t, actual.String(), expected)
	})
	t.Run("pop multiple items", func(t *testing.T) {
		expected := "[]"
		data := bytes.NewBuffer([]byte("[\"One task\",\"Two tasks\"]"))
		var task_list = tasks.LoadState(data)
		task_list, _ = tasks.Pop(task_list)
		task_list, _ = tasks.Pop(task_list)
		task_list, _ = tasks.Pop(task_list)
		actual := tasks.SaveState(task_list)
		test_utils.AssertEqual(t, actual.String(), expected)
	})
}

func TestWorkflowSet(t *testing.T) {
	t.Run("set first with empty list", func(t *testing.T) {
		expected := "[]"
		data := bytes.NewBuffer([]byte("[]"))
		var task_list = tasks.LoadState(data)
		task_list = tasks.Set(task_list, 0, "Task update")
		actual := tasks.SaveState(task_list)
		test_utils.AssertEqual(t, actual.String(), expected)
	})
	t.Run("set first", func(t *testing.T) {
		expected := "[\"Task update\"]"
		data := bytes.NewBuffer([]byte("[\"First task\"]"))
		var task_list = tasks.LoadState(data)
		task_list = tasks.Set(task_list, 1, "Task update")
		actual := tasks.SaveState(task_list)
		test_utils.AssertEqual(t, actual.String(), expected)
	})
	t.Run("set last with existing item", func(t *testing.T) {
		expected := "[\"Task 1\",\"Task 2\",\"Task 3\",\"Task 4\",\"Task 5\",\"Task 6\",\"Task 7\",\"Task 8\",\"Task 9\",\"Task update\"]"
		data := bytes.NewBuffer([]byte(FULL_TASKS))
		var task_list = tasks.LoadState(data)
		task_list = tasks.Set(task_list, 10, "Task update")
		actual := tasks.SaveState(task_list)
		test_utils.AssertEqual(t, actual.String(), expected)
	})
	t.Run("set last without existing item", func(t *testing.T) {
		expected := "[\"Task 1\",\"Task 2\",\"Task 3\",\"Task 4\",\"Task 5\",\"Task 6\",\"Task 7\",\"Task 8\",\"Task 9\"]"
		data := bytes.NewBuffer([]byte("[\"Task 1\",\"Task 2\",\"Task 3\",\"Task 4\",\"Task 5\",\"Task 6\",\"Task 7\",\"Task 8\",\"Task 9\"]"))
		var task_list = tasks.LoadState(data)
		task_list = tasks.Set(task_list, 10, "Task update")
		actual := tasks.SaveState(task_list)
		test_utils.AssertEqual(t, actual.String(), expected)
	})
}
