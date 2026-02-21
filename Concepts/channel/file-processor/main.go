package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	taskChan := make(chan int)
	resultChan := make(chan string)

	go worker(taskChan, resultChan)

	go func() {
		for i := 0; i < 5; i++ {
			taskChan <- i
			time.Sleep(time.Millisecond * 500) // Simulate task creation delay
		}
		close(taskChan) // Signal that no more tasks will be sent
	}()

	for result := range resultChan {
		fmt.Println(result)
	}
}

func worker(taskChan <-chan int, resultChan chan<- string) {
	for task := range taskChan {
		result := processTask(task)
		resultChan <- result
	}
	close(resultChan) // Signal that all results have been sent
}

func processTask(task int) string {
	sleepTime := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(sleepTime)

	return fmt.Sprintf("Task %d processed after %v", task, sleepTime)
}
