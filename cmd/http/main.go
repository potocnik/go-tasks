package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	tasks "tasks/pkg/task_list"
	"tasks/pkg/utils"
)

const PORT = 10000

var TaskList = []string{}

func main() {
	input := utils.ReadFile("tasks.json")
	if input.Len() > 0 {
		TaskList = tasks.LoadState(input)
	}
	fmt.Println("Started on http://localhost:" + fmt.Sprintf("%d", PORT) + "/")
	handleRequests()
}

func listTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		writeToResponse(w)
	case "POST":
		task_text, success := utils.GetParameter(r, "name")
		if success {
			TaskList = tasks.Push(TaskList, task_text)
		}
		writeToResponse(w)
	case "PUT":
		text, textSuccess := utils.GetParameter(r, "name")
		positionStr, positionSuccess := utils.GetParameter(r, "position")
		if textSuccess && positionSuccess {
			position, err := strconv.Atoi(positionStr)
			utils.Check(err)
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
