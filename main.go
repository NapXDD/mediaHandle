package main

import (
	"fmt"
	"transferMP4/media"
)

func main() {
	videoId := "aDjLXFAgsmA"

	audio, err := media.GetYoutubeMP3(videoId)
	if(err != nil){
		fmt.Println(err)
	} else {
		fmt.Println(len(audio))
	}
}
