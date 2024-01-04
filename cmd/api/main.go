package main

import (
	"fmt"

	"github.com/hokkung/go-groceries/di"
	"github.com/hokkung/go-groceries/pkg/env"
	srvutil "github.com/hokkung/srv/util"
)

func main() {
	env.MustLoadEnv()

	container, cleanUp, err := di.InitializeApplication()
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		container.Server.Start()
	}()

	srvutil.WaitForSignalToShutdown()

	cleanUp()

	fmt.Println("application gracefully shutdown.")
}
