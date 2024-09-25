package main

import (
	"context"
	"fmt"
	"github.com/hokkung/go-groceries/internal/di"
	"os/signal"
	"syscall"

	_ "github.com/hokkung/go-groceries/docs"
	"github.com/hokkung/go-groceries/pkg/env"
)

//	@title			Go Groceries APIs
//	@version		1.0
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@schemes	http https

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	env.MustLoadEnv()
	container, cleanUp, err := di.InitializeApplication(ctx)
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
