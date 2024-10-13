package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Id          int       `json:'id'`
	Description string    `json:'description'`
	Status      string    `json:'status'`
	CreatedAt   time.Time `json:'created_at'`
	UpdatedAt   time.Time `json:'updated_at'`
}

// Helper function to list all tasks
func ListTasks(tasks []Task) {
	fmt.Println("ID   | Task Description               | Status      | Created At           | Updated At")
	fmt.Println("-----|--------------------------------|-------------|----------------------|----------------------")
	for _, task := range tasks {
		fmt.Printf("%-4d | %-30s | %-11s | %-20s | %-20s\n", task.Id, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
}

// Helper function to list tasks by status
func ListTasksByStatus(tasks []Task, status string) {
	fmt.Println("ID   | Task Description               | Status      | Created At           | Updated At")
	fmt.Println("-----|--------------------------------|-------------|----------------------|----------------------")
	for _, task := range tasks {
		if task.Status == status {
			fmt.Printf("%-4d | %-30s | %-11s | %-20s | %-20s\n", task.Id, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
		}
	}
}

func LoadTasks(fileName string) ([]Task, error) {

	mutex.Lock()
	defer mutex.Unlock()

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

func SaveTask(tasks []Task, filename string) error {
	mutex.Lock()
	defer mutex.Unlock()

	data, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
