package main

import (
	"fmt"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	// get the filepath from cmd args
	filepath := os.Args[1]
	fmt.Println("filepath: ", filepath)

	// prepare the video file
	vc, err := gocv.VideoCaptureFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// create a window
	window := gocv.NewWindow("Hello")
	// craeate a var called frame to store the current frame
	frame := gocv.NewMat()
	// read frames from the video
	for vc.Read(&frame) {
		// show the current frame in the window
		window.IMShow(frame)
		// frame rate = 33
		c := window.WaitKey(33)
		// if Esc key is pressed, exit
		if c == 27 {
			break
		}
	}
}
