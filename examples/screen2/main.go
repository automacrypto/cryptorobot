package main

import (
	"log"

	"github.com/go-vgo/robotgo"
)

func main() {
	test := robotgo.CaptureImg(10, 10, 30, 30)
	if err := robotgo.Save(test, "test.png"); err != nil {
		log.Fatal(err)
	}
}
