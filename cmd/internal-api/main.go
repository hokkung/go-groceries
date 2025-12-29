package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/hokkung/go-groceries/internal/di"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	container, cleanUp, err := di.InitializeInternalApplication(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- container.Server.Start()
	}()

	select {
	case e := <-srvErr:
		fmt.Println("start running server failed", e.Error())
	case <-ctx.Done():
		fmt.Println("start shutting server down")
	}

	cleanUp()
	fmt.Println("gracefully shutdown")
}
