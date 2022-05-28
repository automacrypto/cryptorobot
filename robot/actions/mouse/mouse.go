package mouse

import (
	"github.com/go-vgo/robotgo"
)

func MouseSleep(miliseconds int) {
	robotgo.MouseSleep = miliseconds
}

func Move(x, y int) {
	robotgo.Move(x, y)
}

func MoveClick(x, y int) {
	robotgo.MoveClick(x, y)
}
