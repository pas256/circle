package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

func main() {
	fmt.Println("Hello World")
	time.Sleep(10 * time.Second)
	fmt.Println("beep")

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("This is a log message", zap.Bool("bool", true))
}
