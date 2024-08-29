package main

import (
	"fmt"
	"net/http"
	"path"
	templates "tasks/cmd/web/templates"
	logger "tasks/pkg/logging"
	tasks "tasks/pkg/task_list"
	file "tasks/pkg/utils/files"
)

const HTTP_PORT = 8090

var TaskList = []string{}

func main() {
	logger.SetUpLogging()
	input := file.ReadFile("../data/tasks.json")
	if input.Len() > 0 {
		TaskList = tasks.LoadState(input)
	}
	logger.Debug("Number of loaded tasks", len(TaskList))
	if len(TaskList) == 0 {
		TaskList = tasks.Push(TaskList, "Example task")
	}
	templates.Handle_Template("/", path.Join("templates", "index.html"), TaskList)
	templates.Handle_Template("/tasks", path.Join("templates", "tasks.html"), nil)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Server started on http://localhost:" + fmt.Sprintf("%d", HTTP_PORT) + "/")
	http.ListenAndServe(fmt.Sprintf("localhost:%d", HTTP_PORT), nil)
	defer tasks.SaveState(TaskList)
}
