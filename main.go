package main

import (
	"fmt"

	"github.com/luminous479/Task-CLI/internal/task"
)

func main() {
	fmt.Println("Task CLI")

	task.Hello()
	newTask := task.CreateTask("learn go")
	fmt.Println(newTask.ID)
	fmt.Println(newTask.Title)
	fmt.Println(newTask.Completed)
	newTask.Complete()
	fmt.Println(newTask.Completed)
    newTask.PrintInfo()
}
