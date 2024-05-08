package main

import (
	"fmt"
	"streamer"
)

func main() {
	// define number and jobs
	const numJobs = 1
	const numWorkers = 2

	// create channel for works and results
	notifyChan := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob, numJobs)
	defer close(videoQueue)

	// Get a worker pool.
	wp := streamer.New(videoQueue, numWorkers)
	fmt.Println(wp)

	// Start the worker pool.
	wp.Run()
	fmt.Println("worker pool started. Press enter to continue")
	_, _ = fmt.Scanln()
	// Create 4 videos to send to the worker pool
	video := wp.NewVideo(1, "./input/puppy1.mp4", "./output", "mp4", notifyChan, nil)

	// Send videos to the worker pool.
	videoQueue <- streamer.VideoProcessingJob{Video: video}
	// Print out results.

	for i := 1; i <= numJobs; i++ {
		msg := <-notifyChan

		fmt.Println("i:", i, "msg:", msg)
	}

}
