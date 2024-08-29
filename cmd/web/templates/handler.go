package templates

import (
	"html/template"
	"net/http"
	errors "tasks/pkg/utils/errors"
)

func Handle_Template(url string, templatePath string, data any) {
	http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(templatePath)
		errors.Check(err)
		err = tmpl.Execute(w, data)
		errors.Check(err)
	})
}
