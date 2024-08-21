package request

import (
	"io"
	"net/http"
	"strings"
	error "tasks/pkg/utils/errors"
)

func Read(stream io.ReadCloser) (string, bool) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, stream)
	error.Check(err)
	if err != nil {
		return "", false
	}
	return buf.String(), true
}

func GetParameter(r *http.Request, name string) (string, bool) {
	errForm := r.ParseForm()
	error.Check(errForm)
	if r.Form.Has(name) {
		return r.Form.Get(name), true
	}
	return "", false
}
