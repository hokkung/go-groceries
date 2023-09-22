package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hokkung/go-groceries/di"
)

func main() {
	container, cleanUp, err := di.InitializeApplication()
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		container.Server.Start()
	}()

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGTERM, syscall.SIGINT)
	<-gracefulShutdown

	cleanUp()

	fmt.Println("application gracefully shutdown.")
}
