package main

import (
	"fmt"
	"echo-go/internal/server"
	"echo-go/internal/utils"
)

func main() {
	fmt.Println("Starting Echo Server")

	fmt.Println("Setting up logger")
	utils.SetupLogger()
	utils.Log.Info("Logger setup complete")

	server.StartServer(":1337")
}
