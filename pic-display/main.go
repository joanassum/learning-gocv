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

	// read the image
	img := gocv.IMRead(filepath, 1)

	// create a window
	window := gocv.NewWindow("Hello")
	// show the image in the window
	window.IMShow(img)
	// wait for 10 sec, then close
	window.WaitKey(10000)
}
