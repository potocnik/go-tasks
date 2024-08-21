package file

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	error "tasks/pkg/utils/errors"
)

func WriteFile(path string, data bytes.Buffer) {
	errDir := os.MkdirAll("data", os.ModePerm)
	error.Check(errDir)
	err := os.WriteFile(fullPath(path), data.Bytes(), 0644)
	error.Check(err)
}

func ReadFile(path string) *bytes.Buffer {
	finalPath := fullPath(path)
	if !fileExists(finalPath) {
		return bytes.NewBuffer([]byte{})
	}
	data, err := os.ReadFile(finalPath)
	error.Check(err)
	return bytes.NewBuffer(data)
}

func fullPath(path string) string {
	return filepath.Join("data", path)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return !errors.Is(err, os.ErrNotExist)
}
