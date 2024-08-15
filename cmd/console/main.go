package main

import (
	"fmt"
	tasks "tasks/pkg/task_list"
	"tasks/pkg/utils"
)

func main() {
	fmt.Println("Writing tasks")
	stream := tasks.WriteTasks()
	utils.Write("tasks.json", stream)
	fmt.Println("Printing tasks")
	lines := tasks.PrintTasks()
	for lineIndex := 0; lineIndex < len(lines); lineIndex++ {
		fmt.Println(lines[lineIndex])
	}
}
