package utils

import (
	"bytes"
	"os"
	"path/filepath"
)

func WriteFile(path string, data bytes.Buffer) {
	errDir := os.MkdirAll("data", os.ModePerm)
	Check(errDir)
	err := os.WriteFile(fullPath(path), data.Bytes(), 0644)
	Check(err)
}

func ReadFile(path string) *bytes.Buffer {
	data, err := os.ReadFile(fullPath((path)))
	Check(err)
	return bytes.NewBuffer(data)
}

func fullPath(path string) string {
	return filepath.Join("data", path)
}
