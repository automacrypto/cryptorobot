package robot

import (
	"fmt"

	"github.com/automacrypto/cryptorobot/robot/helpers"
)

type Robot struct {
	filenames []string
}

func (r *Robot) Start() {
	file := helpers.ReadYaml("robot/steps/01.yaml")
	fmt.Println(file)
}
