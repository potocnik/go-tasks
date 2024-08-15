package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"tasks/pkg/utils"
)

func GetTasks() []string {
	return []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}
}

func PrintTasks() []string {
	var tasks = GetTasks()
	lines := []string{}
	for itemIndex := 0; itemIndex < 10; itemIndex++ {
		task := ""
		if itemIndex < len(tasks) {
			task = tasks[itemIndex]
		}
		lines = append(lines, fmt.Sprintf("%d. %s", itemIndex+1, task))
	}
	return lines
}

func WriteTasks() bytes.Buffer {
	var tasks = GetTasks()
	var memoryStream bytes.Buffer
	jsonData, errMarshal := json.Marshal(tasks)
	utils.CheckWithMessage(errMarshal, "Error marshaling json")
	_, errWrite := memoryStream.WriteString(string(jsonData))
	utils.CheckWithMessage(errWrite, "Error writing to memory stream")
	return memoryStream
}
