package task

import "fmt"

type Task struct{
	ID int
	Title string
	Completed bool
}

func (t *Task) Complete() {
	t.Completed = true 
}


type Manager struct {
	tasks []Task
} 

func (m *Manager) AddTask(title string) Task{

	task := Task{
		ID: len(m.tasks) + 1,
		Title: title,
		Completed: false,
	}
	m.tasks = append(m.tasks, task)
	return task
} 

func (m Manager) ListTasks() {
	
	for _, task := range m.tasks {
		fmt.Println(task)
	}		
	
}