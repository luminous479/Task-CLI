package main

import (
	"fmt"
	"os"
	"strconv"

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
		if len(os.Args) < 3 {
			fmt.Println("Usage: task done <task-id>")
			return
		}
		taskID := os.Args[2]
		id, err := strconv.Atoi(taskID)
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		err = manager.CompleteTask(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Task marked as done")

	case "delete":

		if len(os.Args) < 3 {
			fmt.Println("Usage: task delete <task-id>")
			return
		}

		taskID := os.Args[2]

		id, err := strconv.Atoi(taskID)
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		err = manager.DeleteTask(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Task deleted")

	default:
		fmt.Println("Unknown command")
	}
}
