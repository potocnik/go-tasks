package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	logger "tasks/pkg/logging"
	models "tasks/pkg/models"
	tasks "tasks/pkg/task_list"
	error "tasks/pkg/utils/errors"
	request "tasks/pkg/utils/request"
	"time"
)

func handle_Taks_Get(w http.ResponseWriter) {
	time.Sleep(time.Duration(time.Second * 5))
}

func handle_Tasks_Post(r *http.Request, channel chan models.QueueMessage) {
	task_text, success := request.GetParameter(r, "name")
	if success {
		TaskList = tasks.Push(TaskList, task_text)
		channel <- models.NewQueueMessage(models.HttpOperation_Post, task_text, -1)
	}
}

func handle_Tasks_Put(r *http.Request, channel chan models.QueueMessage) {
	text, textSuccess := request.GetParameter(r, "name")
	positionStr, positionSuccess := request.GetParameter(r, "position")
	if textSuccess && positionSuccess {
		position, err := strconv.Atoi(positionStr)
		error.Check(err)
		tasks.Set(TaskList, position, text)
		channel <- models.NewQueueMessage(models.HttpOperation_Put, text, position)
	}
}

func handle_Tasks_Delete(channel chan models.QueueMessage) {
	TaskList, _ = tasks.Pop(TaskList)
	channel <- models.NewQueueMessage(models.HttpOperation_Delete, "", -1)
}

func handle_Tasks_BadMethod(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Unexpected HTTP method: "+r.Method, http.StatusBadRequest)
}

func writeToResponse(w http.ResponseWriter) {
	logger.Info("Writing to response")
	var result = []string{}
	if len(TaskList) > 0 {
		result = TaskList
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
