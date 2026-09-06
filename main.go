package main

import (
	"fmt"
	"os"

	"github.com/luminous479/Task-CLI/internal/task"
)

func main() {
	fmt.Println("Task CLI")
	manager := task.Manager{}

	if len(os.Args) < 2 {
		fmt.Println("Usage: task <task-title>")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":

		if len(os.Args) < 3 {
			fmt.Println("Usage: task add <title>")
			return
		}

		title := os.Args[2]
		newTask := manager.AddTask(title)

		fmt.Println("Task added:", newTask.Title)

	case "list":
		manager.ListTasks()

	case "done":
		fmt.Println("Complete Task")

	case "delete":
		fmt.Println("Delete Task")

	default:
		fmt.Println("Unknown command")
	}
}
