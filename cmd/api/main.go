package main

import (
	"fmt"

	"github.com/hokkung/go-groceries/di"
)

func main() {
	container, err := di.InitializeApplication()
	if err != nil {
		fmt.Println(err)
		return
	}
	container.Server.Start()
}
