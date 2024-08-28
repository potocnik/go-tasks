package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	models "tasks/pkg/models"
	tasks "tasks/pkg/task_list"
	error "tasks/pkg/utils/errors"
	file "tasks/pkg/utils/files"
)

const PORT = 10000

var TaskList = []string{}
var ChannelWork chan models.QueMessage

func main() {
	setUpLogging()
	setUpChannels()
	input := file.ReadFile("tasks.json")
	if input.Len() > 0 {
		TaskList = tasks.LoadState(input)
	}
	fmt.Println("Started on http://localhost:" + fmt.Sprintf("%d", PORT) + "/")
	handleRequests()
}

func setUpLogging() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	error.Check(err)
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)
}

func setUpChannels() {
	ChannelWork = make(chan models.QueMessage)
}

func handleRequests() {
	http.HandleFunc("/v1/api/tasks", handleTasks)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handle_Taks_Get(w)
	case http.MethodPost:
		handle_Tasks_Post(w, r, ChannelWork)
	case http.MethodPut:
		handle_Tasks_Put(w, r)
	case http.MethodDelete:
		handle_Tasks_Delete(w)
	default:
		handle_Tasks_BadMethod(w, r)
	}
}
