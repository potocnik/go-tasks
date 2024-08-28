package unit__test

import (
	testSubject "tasks/pkg/models"
	"tasks/pkg/test_utils"
	"testing"
)

func TestNewQueMessage(t *testing.T) {
	t.Run("without all arguments creates all data", func(t *testing.T) {
		actual := testSubject.NewQueueMessage(testSubject.HttpOperation_Post, "Task content", 3)
		expected := testSubject.QueueMessage{}
		expected.Operation = testSubject.HttpOperation_Post
		expected.Data = &testSubject.OperationData{}
		expected.Data.Index = 3
		expected.Data.Text = "Task content"
		test_utils.AssertEqual_QueMessage(t, &actual, &expected)
	})
	t.Run("without arguments creates no data", func(t *testing.T) {
		actual := testSubject.NewQueueMessage(testSubject.HttpOperation_Get, "", -1)
		expected := testSubject.QueueMessage{}
		expected.Operation = testSubject.HttpOperation_Get
		test_utils.AssertEqual_QueMessage(t, &actual, &expected)
	})
}
