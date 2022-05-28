package robot

import (
	"fmt"
	"time"

	"github.com/automacrypto/cryptorobot/robot/actions/window"
	"github.com/automacrypto/cryptorobot/robot/helpers"
	"github.com/automacrypto/cryptorobot/robot/models"
)

type Robot struct {
	filenames []string
}

func (r *Robot) Start() {
	config := helpers.ReadYaml[models.Config]("config.yaml")

	for index := range config.Macros {
		macro := helpers.ReadYaml[models.Macro](config.Paths.Macros + config.Macros[index])

		hwnd := window.FindWindow(macro.Name)

		fmt.Println("I found a window", macro.Name, "with PID", hwnd)

		success := window.SetForegroundWindow(hwnd)

		if !success {
			fmt.Println("Fail to set active window")
		}

		const zzz = 60
		fmt.Println("I'm tired, I'm going to sleep for", zzz, "seconds")
		time.Sleep(zzz * time.Second)
	}
}
