package streamer

type ProcessingMessage struct {
	ID         int
	Successful bool
	Message    string
	OutputFile string
}

type VideoProcessingJob struct {
	Video Video
}
type Processor struct {
	Engine Encoder
}

type Video struct {
	ID            int
	InputFile     string
	OutputDir     string
	EncodingType  string
	NotifyChannel chan ProcessingMessage
	Encoder       Processor
}

func New(jobQueue chan VideoProcessingJob, maxWorkers int) *VideoDispatcher {
	workerPool := make(chan chan VideoProcessingJob, maxWorkers)
	// implement processor logic
	p := Processor{}
	return &VideoDispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		WorkerPool: workerPool,
		Processor:  p,
	}
}
