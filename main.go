package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/luminous479/Task-CLI/internal/task"
)

func main() {
	manager := task.Manager{}

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  task add <title>")
		fmt.Println("  task list")
		fmt.Println("  task done <id>")
		fmt.Println("  task delete <id>")
		return
	}

	command := os.Args[1]

	switch command {

	case "add":
		addTask(&manager)

	case "list":
		manager.ListTasks()

	case "done":
		completeTask(&manager)

	case "delete":
		deleteTask(&manager)

	default:
		fmt.Println("Unknown command:", command)
	}
}

func addTask(manager *task.Manager) {
	if len(os.Args) < 3 {
		fmt.Println("Usage: task add <title>")
		return
	}

	title := os.Args[2]

	newTask := manager.AddTask(title)

	fmt.Println("Task added:", newTask.Title)
}

func completeTask(manager *task.Manager) {
	if len(os.Args) < 3 {
		fmt.Println("Usage: task done <id>")
		return
	}

	id, err := strconv.Atoi(os.Args[2])

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
}

func deleteTask(manager *task.Manager) {
	if len(os.Args) < 3 {
		fmt.Println("Usage: task delete <id>")
		return
	}

	id, err := strconv.Atoi(os.Args[2])

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
}