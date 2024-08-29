package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	logger "tasks/pkg/logging"
	array "tasks/pkg/utils/arrays"
	error "tasks/pkg/utils/errors"
)

const ITEM_LIMIT = 10

func Print(task_list []string) []string {
	lines := []string{}
	for itemIndex := 0; itemIndex < ITEM_LIMIT; itemIndex++ {
		task := ""
		if itemIndex < len(task_list) {
			task = task_list[itemIndex]
		}
		lines = append(lines, fmt.Sprintf("%d. %s", itemIndex+1, task))
	}
	return lines
}

func Push(task_list []string, task_text string) []string {
	if task_text == "" {
		logger.Error("cannot append empty task", nil)
	} else if len(task_list) >= ITEM_LIMIT {
		logger.Error("list is full (limit "+fmt.Sprintf("%d", ITEM_LIMIT)+")", nil)
	} else {
		logger.Debug("Appending task", task_text)
		task_list = append(task_list, task_text)
	}
	return task_list
}

func Pop(task_list []string) ([]string, *string) {
	if len(task_list) < 1 {
		return task_list, nil
	}
	task := task_list[0]
	task_list = array.RemoveAt(task_list, 0)
	return task_list, &task
}

func Set(task_list []string, position int, task_text string) []string {
	if position > 0 && len(task_list) >= position {
		task_list[position-1] = task_text
	}
	return task_list
}

func LoadState(stream *bytes.Buffer) []string {
	return readTasks(stream)
}

func SaveState(task_list []string) bytes.Buffer {
	logger.Info("[INFO] Writing tasks to data/tasks.json")
	stream := writeTasks(task_list)
	return stream
}

func writeTasks(task_list []string) bytes.Buffer {
	var memoryStream bytes.Buffer
	if len(task_list) == 0 {
		_, errWrite := memoryStream.WriteString("[]")
		error.CheckWithMessage(errWrite, "Error writing to memory stream")
	} else {
		jsonData, errMarshal := json.Marshal(task_list)
		error.CheckWithMessage(errMarshal, "Error marshaling json")
		_, errWrite := memoryStream.WriteString(string(jsonData))
		error.CheckWithMessage(errWrite, "Error writing to memory stream")
	}
	return memoryStream
}

func readTasks(stream *bytes.Buffer) []string {
	var data []string
	err := json.Unmarshal(stream.Bytes(), &data)
	error.Check(err)
	return data
}
