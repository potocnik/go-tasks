package main

import "fmt"

func main() {
	fmt.Println("Starting")
	lines := PrintTasks()
	for lineIndex := 0; lineIndex < len(lines); lineIndex++ {
		fmt.Println(lines[lineIndex])
	}
}
