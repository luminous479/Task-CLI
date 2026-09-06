package task

import "fmt"

type Task struct{
	ID int
	Title string
	Completed bool
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
func (m *Manager) CompleteTask(id int) error {
    for i, task := range m.tasks {
        if task.ID == id {
            m.tasks[i].Completed = true
            return nil
        }
    }

    return fmt.Errorf("task with ID %d not found", id)
}

func (m *Manager) DeleteTask(id int) error {
	for i, task := range m.tasks {
		if task.ID == id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}
 
