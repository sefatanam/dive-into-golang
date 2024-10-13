package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type Task struct {
	Id          int       `json:'id'`
	Description string    `json:'description'`
	Status      string    `json:'status'`
	CreatedAt   time.Time `json:'created_at'`
	UpdatedAt   time.Time `json:'updated_at'`
}

// global variables
const (
	fileName = "tasks.json"
)

var mutex sync.Mutex

func loadTasks(fileName string) ([]Task, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []Task{}, err
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return []Task{}, err
	}
	var tasks []Task

	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

// Helper function to list all tasks
func listTasks(tasks []Task) {
	fmt.Println("ID   | Task Description               | Status      | Created At           | Updated At")
	fmt.Println("-----|--------------------------------|-------------|----------------------|----------------------")
	for _, task := range tasks {
		fmt.Printf("%-4d | %-30s | %-11s | %-20s | %-20s\n", task.Id, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
}

// Helper function to list tasks by status
func listTasksByStatus(tasks []Task, status string) {
	fmt.Println("ID   | Task Description               | Status      | Created At           | Updated At")
	fmt.Println("-----|--------------------------------|-------------|----------------------|----------------------")
	for _, task := range tasks {
		if task.Status == status {
			fmt.Printf("%-4d | %-30s | %-11s | %-20s | %-20s\n", task.Id, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli [add|update|delete|mark-in-progess|mark-done|list] [task details]")
		return
	}

	command := os.Args[1]

	tasks, err := loadTasks(fileName)

	if err != nil {
		fmt.Println("Error loading tasks: ", err)
		return
	}

	switch command {
	case "list":
		if len(os.Args) > 2 {
			status := os.Args[2]
			listTasksByStatus(tasks, status)
		} else {
			listTasks(tasks)
		}

	default:
		fmt.Println("Unknown command:", command)
	}

}
