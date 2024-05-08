package streamer

import (
	"fmt"

	"github.com/xfrr/goffmpeg/transcoder"
)

type Encoder interface {
	EncodeToMP4(v *Video, baseFileName string) error
}

type VideoEncoder struct{}

func (ve *VideoEncoder) EncodeToMP4(v *Video, baseFileName string) error {
	// create a transcoder
	trans := new(transcoder.Transcoder)
	// output path
	outputPath := fmt.Sprintf("%s/%s.mp4", v.OutputDir, baseFileName)
	// initialise transocder
	err := trans.Initialize(v.InputFile, outputPath)
	if err != nil {
		return err
	}
	// set codec
	trans.MediaFile().SetVideoCodec("libx264")
	// start the transcoding process
	done := trans.Run(false)
	err = <-done
	if err != nil {
		return err
	}
	return nil
}
