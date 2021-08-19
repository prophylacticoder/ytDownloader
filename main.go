package main

import (
  "fmt"
  "os"

  "github.com/kkdai/youtube/v2"
)

func main() {
  // Checks if the URL was inserted
  if len(os.Args) < 2 {
    fmt.Fprintf(os.Stderr, "usage: %s [video ID]", os.Args[0])
    return
  }
  // Creates a client
  client := youtube.Client{}
  // Gets the video
  video, err := cleint.GetVideo(video)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error getting the video.")
    return
  }

  stream, _, err := client.GetStream(video, &video.Formats[0])
	if err != nil {
		panic(err)
	}

  file, err := os.Create("video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}
