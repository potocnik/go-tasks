package main

import (
	"fmt"
	"os"
	logger "tasks/pkg/logging"
	tasks "tasks/pkg/task_list"
	file "tasks/pkg/utils/files"
)

func main() {
	logger.SetUp()
	logger.Debug("Reading tasks from data/tasks.json", nil)
	input := file.ReadFile("tasks.json")
	var task_list = tasks.LoadState(input)
	task_list = processCommand(task_list)
	logger.Debug("Printing tasks to screen", nil)
	lines := tasks.Print(task_list)
	for lineIndex := 0; lineIndex < len(lines); lineIndex++ {
		fmt.Println(lines[lineIndex])
	}
	stream := tasks.SaveState(task_list)
	logger.Debug("Writing tasks to data/tasks.json", nil)
	file.WriteFile("tasks.json", stream)
}

func processCommand(task_list []string) []string {
	logger.Info("Processing commands")
	if len(os.Args) > 1 {
		var command = os.Args[1]
		switch command {
		case "push":
			var commandArg = ""
			if len(os.Args) > 3 {
				commandArg = os.Args[2]
			}
			return tasks.Push(task_list, commandArg)
		case "pop":
			tasks, _ := tasks.Pop(task_list)
			return tasks
		}
	}
	return task_list
}
