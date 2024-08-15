package test_utils

import (
	"reflect"
	"testing"
)

func AssertEqual(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %q want %q", actual, expected)
	}
}

func AssertEqualArray(t testing.TB, actual, expected []string) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}
