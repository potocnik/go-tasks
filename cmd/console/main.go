package main

import (
	"fmt"
	"os"
	tasks "tasks/pkg/task_list"
	"tasks/pkg/utils"
)

func main() {
	var task_list = loadState()
	task_list = processCommand(task_list)
	fmt.Println("[INFO] Printing tasks to screen")
	lines := tasks.PrintTasks(task_list)
	for lineIndex := 0; lineIndex < len(lines); lineIndex++ {
		fmt.Println(lines[lineIndex])
	}
	saveState(task_list)
}

func processCommand(task_list []string) []string {
	fmt.Println("[INFO] Processing commands")
	if len(os.Args) > 1 {
		var command = os.Args[1]
		switch command {
		case "add":
			return processCommandAdd(task_list)
		}
	}
	return task_list
}

func processCommandAdd(task_list []string) []string {
	if len(os.Args) <= 2 {
		fmt.Println("[ERROR] command add requires an additional argument")
	} else {
		var taskText = os.Args[2]
		task_list = append(task_list, taskText)
	}
	return task_list
}

func loadState() []string {
	fmt.Println("[INFO] Reading tasks from data/tasks.json")
	input := utils.ReadFile("tasks.json")
	return tasks.ReadTasks(input)
}

func saveState(task_list []string) {
	fmt.Println("[INFO] Writing tasks to data/tasks.json")
	stream := tasks.WriteTasks(task_list)
	fmt.Print("[INFO] JSON: ")
	fmt.Println(stream.String())
	utils.WriteFile("tasks.json", stream)
}
