package streamer

type VideoDispatcher struct {
	WorkerPool chan chan VideoProcessingJob
	maxWorkers int
	jobQueue   chan VideoProcessingJob
	Processor  Processor
}

type videoWorker struct {
	ID         int
	jobQueue   chan VideoProcessingJob
	workerPool chan chan VideoProcessingJob
}

func (vd *VideoDispatcher) Run() {
	for i := 0; i < vd.maxWorkers; i++ {
		worker := newVideoWorker(i+1, vd.WorkerPool)
		worker.start()
	}
	go vd.dispatch()
}

func newVideoWorker(id int, workerPool chan chan VideoProcessingJob) videoWorker {
	return videoWorker{
		ID:         id,
		jobQueue:   make(chan VideoProcessingJob),
		workerPool: workerPool,
	}
}

func (w videoWorker) start() {
	go func() {
		for {
			w.workerPool <- w.jobQueue
			job := <-w.jobQueue
			w.processVideoJob(job.Video)
		}
	}()
}

func (vd *VideoDispatcher) dispatch() {
	for {
		job := <-vd.jobQueue

		go func() {
			workerJobQueue := <-vd.WorkerPool
			workerJobQueue <- job
		}()
	}
}

func (w videoWorker) processVideoJob(video Video) {
	video.encode()
}
