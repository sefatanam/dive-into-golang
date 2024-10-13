package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// global variables
const (
	fileName = "tasks.json"
)

var mutex sync.Mutex

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli [add|update|delete|mark-in-progess|mark-done|list] [task details]")
		return
	}

	command := os.Args[1]

	tasks, err := LoadTasks(fileName)

	if err != nil {
		fmt.Println("Error loading tasks: ", err)
		return
	}

	switch command {
	case "list":
		if len(os.Args) > 2 {
			status := os.Args[2]
			ListTasksByStatus(tasks, status)
		} else {
			ListTasks(tasks)
		}
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add [task description]")
			return
		}

		description := os.Args[2]
		id := len(tasks) + 1

		task := Task{
			Id:          id,
			Description: description,
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		tasks = append(tasks, task)
		SaveTask(tasks, fileName)
		fmt.Println("Task added successfully (Id:", id, ")")

	case "update":
		if len(os.Args) > 4 {
			fmt.Println("Usage: task-cli update [id] [desciption]")
			return
		}

		id, _ := strconv.Atoi(os.Args[2])
		taskIndex := GetTaskIndexById(tasks, id)

		if taskIndex == -1 {
			fmt.Println("Task with the ID: ", id, " not found.")
			return
		}

		newDescription := os.Args[3]
		tasks[taskIndex].Description = newDescription
		tasks[taskIndex].UpdatedAt = time.Now()

		SaveTask(tasks, fileName)
		fmt.Println("Task updated successfully (Id :)", id, ")")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage task-cli delete [id]")
			return
		}

		id, _ := strconv.Atoi(os.Args[2])
		taskIndex := GetTaskIndexById(tasks, id)
		if taskIndex == -1 {
			fmt.Println("Task with the ID: ", id, " not found.")
			return
		}

		tasks = append(tasks[:taskIndex], tasks[taskIndex+1:]...)
		SaveTask(tasks, fileName)
		fmt.Println("Task deleted successfully (Id :)", id, ")")
	default:
		fmt.Println("Unknown command:", command)
	}

}
