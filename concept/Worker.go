package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	Id int
}

type Worker struct {
	Id        int
	JobQueue  chan Job
	QuitChan  chan struct{}
	WaitGroup *sync.WaitGroup
}

func NewWorker(id int, jobQueue chan Job, wg *sync.WaitGroup) *Worker {

	return &Worker{
		Id:        id,
		JobQueue:  jobQueue,
		WaitGroup: wg,
		QuitChan:  make(chan struct{}),
	}
}

func (w *Worker) Start() {
	go func() {

		defer w.WaitGroup.Done()
		for {
			select {
			case job := <-w.JobQueue:
				{
					fmt.Printf("Worker %d started job %d\n", w.Id, job.Id)
					time.Sleep(2 * time.Second)
					fmt.Printf("Worker %d finished job %d\n", w.Id, job.Id)
				}

			case <-w.QuitChan:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- struct{}{}
	}()
}

type Pool struct {
	Workers   []*Worker
	JobQueue  chan Job
	WaitGroup *sync.WaitGroup
}

func NewPool(numWorkers int) *Pool {
	return &Pool{
		Workers:   make([]*Worker, numWorkers),
		JobQueue:  make(chan Job),
		WaitGroup: &sync.WaitGroup{},
	}
}

func (p *Pool) Start() {
	for i := 0; i < len(p.Workers); i++ {
		p.Workers[i] = NewWorker(i, p.JobQueue, p.WaitGroup)
		p.Workers[i].Start()
		p.WaitGroup.Add(1)
	}
}

func (p *Pool) SubmitJob(job Job) {
	p.JobQueue <- job
}

func (p *Pool) Stop() {
	close(p.JobQueue)
	p.WaitGroup.Wait()
	for _, worker := range p.Workers {
		worker.Stop()
	}
}

func WorkerPool() {

	pool := NewPool(3)
	pool.Start()

	for i := 0; i < 10; i++ {
		job := Job{Id: i}
		pool.SubmitJob(job)
	}

	pool.Start()
}
