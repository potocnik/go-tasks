package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"tasks/pkg/utils"
)

func PrintTasks(task_list []string) []string {
	lines := []string{}
	for itemIndex := 0; itemIndex < 10; itemIndex++ {
		task := ""
		if itemIndex < len(task_list) {
			task = task_list[itemIndex]
		}
		lines = append(lines, fmt.Sprintf("%d. %s", itemIndex+1, task))
	}
	return lines
}

func WriteTasks(task_list []string) bytes.Buffer {
	var memoryStream bytes.Buffer
	jsonData, errMarshal := json.Marshal(task_list)
	utils.CheckWithMessage(errMarshal, "Error marshaling json")
	_, errWrite := memoryStream.WriteString(string(jsonData))
	utils.CheckWithMessage(errWrite, "Error writing to memory stream")
	return memoryStream
}

func ReadTasks(stream *bytes.Buffer) []string {
	var data []string
	err := json.Unmarshal(stream.Bytes(), &data)
	utils.Check(err)
	return data
}
