package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	models "tasks/pkg/models"
	tasks "tasks/pkg/task_list"
	error "tasks/pkg/utils/errors"
	request "tasks/pkg/utils/request"
)

func handle_Taks_Get(w http.ResponseWriter) {
	writeToResponse(w)
}

func handle_Tasks_Post(w http.ResponseWriter, r *http.Request, channel chan models.QueMessage) {
	task_text, success := request.GetParameter(r, "name")
	if success {
		TaskList = tasks.Push(TaskList, task_text)
	}
	channel <- models.NewQueueMessage(models.HttpOperation_Get, task_text, -1)
	writeToResponse(w)
}

func handle_Tasks_Put(w http.ResponseWriter, r *http.Request) {
	text, textSuccess := request.GetParameter(r, "name")
	positionStr, positionSuccess := request.GetParameter(r, "position")
	if textSuccess && positionSuccess {
		position, err := strconv.Atoi(positionStr)
		error.Check(err)
		tasks.Set(TaskList, position, text)
	}
	writeToResponse(w)
}

func handle_Tasks_Delete(w http.ResponseWriter) {
	TaskList, _ = tasks.Pop(TaskList)
	writeToResponse(w)
}

func handle_Tasks_BadMethod(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Unexpected HTTP method: "+r.Method, http.StatusBadRequest)
}

func writeToResponse(w http.ResponseWriter) {
	var result = []string{}
	if len(TaskList) > 0 {
		result = TaskList
	}
	json.NewEncoder(w).Encode(result)
}
