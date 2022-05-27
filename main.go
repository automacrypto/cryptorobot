package main

import (
	"fmt"

	"github.com/automacrypto/cryptorobot/robot"
)

func main() {
	fmt.Println("Hi human")

	robot := robot.Robot{}

	robot.Start()
}
