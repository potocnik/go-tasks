package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	tasks "tasks/pkg/task_list"
	error "tasks/pkg/utils/errors"
	file "tasks/pkg/utils/files"
	request "tasks/pkg/utils/request"
)

const PORT = 10000

var TaskList = []string{}

func main() {
	setUpLogging()
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

func listTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		writeToResponse(w)
	case "POST":
		task_text, success := request.GetParameter(r, "name")
		if success {
			TaskList = tasks.Push(TaskList, task_text)
		}
		writeToResponse(w)
	case "PUT":
		text, textSuccess := request.GetParameter(r, "name")
		positionStr, positionSuccess := request.GetParameter(r, "position")
		if textSuccess && positionSuccess {
			position, err := strconv.Atoi(positionStr)
			error.Check(err)
			tasks.Set(TaskList, position, text)
		}
		writeToResponse(w)
	case "DELETE":
		TaskList, _ = tasks.Pop(TaskList)
		writeToResponse(w)
	default:
		http.Error(w, "Unexpected HTTP method: "+r.Method, http.StatusBadRequest)
	}
}

func writeToResponse(w http.ResponseWriter) {
	var result = []string{}
	if len(TaskList) > 0 {
		result = TaskList
	}
	json.NewEncoder(w).Encode(result)
}

func handleRequests() {
	http.HandleFunc("/v1/api/tasks", listTasks)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
