package streamer

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

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
	Options       *VideoOptions
}

type VideoOptions struct {
	RenameOutput    bool
	SegmentDuration int
	MaxRate1080p    string
	MaxRate720p     string
	MaxRate480p     string
}

func (vd *VideoDispatcher) NewVideo(id int, input, output, encType string, notifyChan chan ProcessingMessage, ops *VideoOptions) Video {
	if ops == nil {
		ops = &VideoOptions{}
	}
	fmt.Println("c.newVideo():new video created", id, input)
	return Video{
		ID:            id,
		InputFile:     input,
		OutputDir:     output,
		EncodingType:  encType,
		NotifyChannel: notifyChan,
		Encoder:       vd.Processor,
		Options:       ops,
	}
}

func (v *Video) encode() {
	var fileName string
	switch v.EncodingType {

	case "mp4":
		fmt.Println("v.encode():about to encde to mp4", v.ID)

		name, err := v.encodeToMp4()
		if err != nil {
			v.sendToNotifyChan(false, "", fmt.Sprintf("encode failed for %d, %s", v.ID, err.Error()))
			return
		}
		fileName = fmt.Sprintf("%s.mp4", name)
	default:
		fmt.Println("v.encode(): error trying to encode video", v.ID)

		v.sendToNotifyChan(false, "", fmt.Sprintf("error processing for %d invalid encoding type", v.ID))
		return
	}
	fmt.Println("v.encode(): sending success message for video id", v.ID, "to notify channel")

	v.sendToNotifyChan(true, fileName, fmt.Sprintf("video id  %d processed and saved as %s", v.ID, fmt.Sprintf("%s/%s", v.OutputDir, fileName)))
}

func (v *Video) encodeToMp4() (string, error) {
	baseFileName := ""
	fmt.Println("v.encodeToMp4(): about to try encode video", v.ID)

	if !v.Options.RenameOutput {
		// get base file name
		b := path.Base(v.InputFile)
		baseFileName = strings.TrimSuffix(b, filepath.Ext(b))
	}

	err := v.Encoder.Engine.EncodeToMP4(v, baseFileName)
	if err != nil {
		return "", err
	}
	fmt.Println("v.encodeToMp4(): successfully encoded", v.ID)

	return baseFileName, nil
}

func (v *Video) sendToNotifyChan(successful bool, fileName, message string) {
	v.NotifyChannel <- ProcessingMessage{
		ID:         v.ID,
		Successful: successful,
		Message:    message,
		OutputFile: fileName,
	}
}

func New(jobQueue chan VideoProcessingJob, maxWorkers int) *VideoDispatcher {
	workerPool := make(chan chan VideoProcessingJob, maxWorkers)

	var e VideoEncoder
	p := Processor{
		Engine: &e,
	}
	return &VideoDispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		WorkerPool: workerPool,
		Processor:  p,
	}
}
