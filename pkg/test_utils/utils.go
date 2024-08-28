package test_utils

import (
	"reflect"
	models "tasks/pkg/models"
	"testing"
)

func AssertEqual_String(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %q want %q", actual, expected)
	}
}

func AssertEqual_Int(t testing.TB, actual, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %q want %q", actual, expected)
	}
}

func AssertEqual_Array(t testing.TB, actual, expected []string) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func AssertEqual_QueMessage(t testing.TB, actual *models.QueMessage, expected *models.QueMessage) {
	t.Helper()
	if actual.Operation != expected.Operation {
		t.Errorf("got %v want %v", actual.Operation, expected.Operation)
	}
	if actual.Data == nil && expected.Data == nil {
		return
	}
	if actual.Data.Index != expected.Data.Index {
		t.Errorf("got %v want %v", actual.Data.Index, expected.Data.Index)
	}
	if actual.Data.Text != expected.Data.Text {
		t.Errorf("got %v want %v", actual.Data.Text, expected.Data.Text)
	}
}
