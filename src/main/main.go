package main

import (
	"fmt"

	"lib/logger"

	_ "lib/models"
)

func main() {
	fmt.Println("Hello World !")
	logger.Info("INFO : This is info message")
	logger.Warn("WARN : This is warn message")
}
