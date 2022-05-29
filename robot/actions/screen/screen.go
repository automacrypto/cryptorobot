package screen

import (
	"image"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/gcv"
	"github.com/vova616/screenshot"
)

func SaveCapture(path string) string {
	capture, _ := screenshot.CaptureScreen()
	printscreen := image.Image(capture)
	now := time.Now().Format("01_02_2006_15:04:05")

	filename := "printscreen_" + now + ".png"
	fullpath := path + "/" + filename

	if error := robotgo.Save(printscreen, fullpath); error != nil {
		panic(error)
	}

	return filename
}

func FindImageFile(fileSearch string, file string) {
	gcv.FindAllImgFile(fileSearch, file)
}
