package file

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	error "tasks/pkg/utils/errors"
)

func WriteFile(path string, data bytes.Buffer) {
	errDir := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	error.Check(errDir)
	err := os.WriteFile(path, data.Bytes(), 0644)
	error.Check(err)
}

func ReadFile(path string) *bytes.Buffer {
	if !fileExists(path) {
		return bytes.NewBuffer([]byte{})
	}
	data, err := os.ReadFile(path)
	error.Check(err)
	return bytes.NewBuffer(data)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return !errors.Is(err, os.ErrNotExist)
}
