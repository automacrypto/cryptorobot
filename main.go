package main

import (
	"fmt"

	"github.com/automacrypto/cryptorobot/app"
)

func main() {
	fmt.Println("Hi human")

	agent := app.Agent{
		name: "CryptoRobot",
	}

	agent.start()
}
