package main

import (
	"abank/app"
	"abank/logger"
)

func main() {
	logger.Info("Starting the application..")
	app.Start()
}
