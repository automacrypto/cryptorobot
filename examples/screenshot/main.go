package main

import (
	"image"
	"log"

	"github.com/go-vgo/robotgo"
	"github.com/vova616/screenshot"
)

func main() {
	img, _ := screenshot.CaptureScreen()
	myImg := image.Image(img)
	if err := robotgo.Save(myImg, "test.png"); err != nil {
		log.Fatal(err)
	}
}
