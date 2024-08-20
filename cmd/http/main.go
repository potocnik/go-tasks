package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	tasks "tasks/pkg/task_list"
	"tasks/pkg/utils"
)

const PORT = 10000

var TaskList = []string{}

func listTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		task_text := utils.Read(r.Body)
		TaskList = tasks.Push(TaskList, task_text)
	case "DELETE":
		TaskList, _ = tasks.Pop(TaskList)
	}
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

func main() {
	input := utils.ReadFile("tasks.json")
	if input.Len() > 0 {
		TaskList = tasks.LoadState(input)
	}
	fmt.Println("Started on http://localhost:" + fmt.Sprintf("%d", PORT) + "/")
	handleRequests()
}
