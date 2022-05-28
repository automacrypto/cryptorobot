package window

import (
	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
)

func FindWindow(name string) win.HWND {
	hwnd := robotgo.FindWindow(name)
	return hwnd
}

func SetActiveWindow(hwnd win.HWND) {
	robotgo.SetActiveWindow(hwnd)
}

func SetFocus(hwnd win.HWND) win.HWND {
	return robotgo.SetFocus(hwnd)
}

func SetForegroundWindow(hwnd win.HWND) bool {
	success := win.SetForegroundWindow(hwnd)
	return success
}
