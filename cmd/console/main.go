package main

import (
	"fmt"
	"os"
	tasks "tasks/pkg/task_list"
	"tasks/pkg/utils"
)

func main() {
	fmt.Println("[DEBUG] Reading tasks from data/tasks.json")
	input := utils.ReadFile("tasks.json")
	var task_list = tasks.LoadState(input)
	task_list = processCommand(task_list)
	fmt.Println("[DEBUG] Printing tasks to screen")
	lines := tasks.Print(task_list)
	for lineIndex := 0; lineIndex < len(lines); lineIndex++ {
		fmt.Println(lines[lineIndex])
	}
	stream := tasks.SaveState(task_list)
	fmt.Println("[DEBUG] Writing tasks to data/tasks.json")
	utils.WriteFile("tasks.json", stream)
}

func processCommand(task_list []string) []string {
	fmt.Println("[INFO] Processing commands")
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
