package main

import (
	"fmt"
	"net/http"
	"path"
	templates "tasks/cmd/web/templates"
	logger "tasks/pkg/logging"
)

const HTTP_PORT = 8090

func main() {
	templates.Handle_Template("/", path.Join("templates", "index.html"))
	templates.Handle_Template("/tasks", path.Join("templates", "tasks.html"))
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	logger.Info("Server started on http://localhost:" + fmt.Sprintf("%d", HTTP_PORT) + "/")
	http.ListenAndServe(fmt.Sprintf("localhost:%d", HTTP_PORT), nil)
}
