package utils

import (
	"io"
	"net/http"
	"strings"
)

func Read(stream io.ReadCloser) (string, bool) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, stream)
	Check(err)
	if err != nil {
		return "", false
	}
	return buf.String(), true
}

func GetParameter(r *http.Request, name string) (string, bool) {
	errForm := r.ParseForm()
	Check(errForm)
	if r.Form.Has(name) {
		return r.Form.Get(name), true
	}
	return "", false
}
