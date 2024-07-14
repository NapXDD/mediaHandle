package main

import (
	"fmt"
	"os"
	"transferMP4/media"
)

func main() {
	videoId := "aDjLXFAgsmA"
	outputFile := "output/" + videoId + ".mp3"
	inputFile := "mp4source/" + videoId + ".mp4"

	media.GetYoutubeVideo(videoId)
	err := media.ConvertMp4ToMp3(inputFile, outputFile)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Conversion successful!")
		os.Remove(inputFile)
	}
}
