package task

import "fmt"

func Hello() {
    fmt.Println("Hello from task package")
}

type Task struct{
	ID int
	Title string
	Completed bool
}


func CreateTask(title string) Task {
    return Task{
        ID:        1,
        Title:     title,
        Completed: false,
    }
}

func (t *Task) Complete() {
	t.Completed = true 
}

func (t Task) PrintInfo(){
	fmt.Println(t.ID)
	fmt.Println(t.Title)
	fmt.Println(t.Completed)
}