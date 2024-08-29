package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	logger "tasks/pkg/logging"
	models "tasks/pkg/models"
	tasks "tasks/pkg/task_list"
	error "tasks/pkg/utils/errors"
	file "tasks/pkg/utils/files"
	"time"
)

const PORT = 10000
const REQUEST_TIMEOUT = time.Minute * 20

var TaskList = []string{}
var channelWork chan models.QueueMessage

func main() {
	setUpLogging()
	setUpChannels()
	input := file.ReadFile("../data/tasks.json")
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
	channelWork = make(chan models.QueueMessage)
}

func handleRequests() {
	http.HandleFunc("/v1/api/tasks", handleTasks)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", PORT), nil))
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	var group sync.WaitGroup
	switch r.Method {
	case http.MethodGet:
		handle_Taks_Get(w)
	case http.MethodPost:
		logger.Info("Starting POST /tasks")
		group.Add(1)
		go handle_Tasks_Post(r, channelWork)
		go handleQueueMessage(channelWork, &group)
	case http.MethodPut:
		logger.Info("Starting PUT /tasks")
		group.Add(1)
		go handle_Tasks_Put(r, channelWork)
		go handleQueueMessage(channelWork, &group)
	case http.MethodDelete:
		logger.Info("Starting DELETE /tasks")
		group.Add(1)
		go handle_Tasks_Delete(channelWork)
		go handleQueueMessage(channelWork, &group)
	default:
		handle_Tasks_BadMethod(w, r)
	}
	group.Wait()
	writeToResponse(w)
}

func handleQueueMessage(channel chan models.QueueMessage, group *sync.WaitGroup) {
	select {
	case queueMessage := <-channel:
		logger.Debug("Consumer Received:", queueMessage)
		group.Done()
	case <-time.After(time.Second * 2):
		// #TODO: Return timeout status code
		group.Done()
	}
}
