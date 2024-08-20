package integration_test

import (
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

// func TestWorkflow_PushWithEmpty(t *testing.T) {
// 	t.Run("go run . push \"First task\"", func(t *testing.T) {
// 		task_list := []string{}
// 		var actual = tasks.(task_list)
// 	})
// }
