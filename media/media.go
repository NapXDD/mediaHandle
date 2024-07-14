package media

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/kkdai/youtube/v2"
)

func ConvertMp4ToMp3(inputFile string, outputFile string) error {
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist")
	}

	cmd := exec.Command("ffmpeg", "-i", inputFile, "-q:a", "0", "-map", "a", outputFile)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error converting file: %v", err)
	}

	return nil
}

func GetYoutubeMP3(videoID string) ([]byte, error) {
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	audioData, err := io.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	return audioData, nil
}
