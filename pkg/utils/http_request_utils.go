package utils

import (
	"io"
	"strings"
)

func Read(stream io.ReadCloser) string {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, stream)
	Check(err)
	return buf.String()
}
