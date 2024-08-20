package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"tasks/pkg/utils"
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
		fmt.Println("[ERROR]: cannot append empty task")
	} else if len(task_list) >= ITEM_LIMIT {
		fmt.Println("[ERROR]: list is full (limit " + fmt.Sprintf("%d", ITEM_LIMIT) + ")")
	} else {
		fmt.Println("[DEBUG]: Appending task: \"" + task_text + "\"")
		task_list = append(task_list, task_text)
	}
	return task_list
}

func Pop(task_list []string) ([]string, *string) {
	if len(task_list) < 1 {
		return task_list, nil
	}
	task := task_list[0]
	task_list = utils.RemoveAt(task_list, 0)
	return task_list, &task
}

func LoadState(stream *bytes.Buffer) []string {
	return readTasks(stream)
}

func SaveState(task_list []string) bytes.Buffer {
	fmt.Println("[INFO] Writing tasks to data/tasks.json")
	stream := writeTasks(task_list)
	return stream
}

func writeTasks(task_list []string) bytes.Buffer {
	var memoryStream bytes.Buffer
	jsonData, errMarshal := json.Marshal(task_list)
	utils.CheckWithMessage(errMarshal, "Error marshaling json")
	_, errWrite := memoryStream.WriteString(string(jsonData))
	utils.CheckWithMessage(errWrite, "Error writing to memory stream")
	return memoryStream
}

func readTasks(stream *bytes.Buffer) []string {
	var data []string
	err := json.Unmarshal(stream.Bytes(), &data)
	utils.Check(err)
	return data
}
