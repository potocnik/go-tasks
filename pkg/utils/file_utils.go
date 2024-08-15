package utils

import (
	"bytes"
	"os"
	"path/filepath"
)

func Write(path string, data bytes.Buffer) {
	errDir := os.MkdirAll("data", os.ModePerm)
	if errDir != nil {
		panic(errDir)
	}
	fullPath := filepath.Join("data", path)
	err := os.WriteFile(fullPath, data.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
