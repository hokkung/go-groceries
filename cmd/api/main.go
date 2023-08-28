package main

import (
	"fmt"

	"github.com/hokkung/go-groceries/di"
)

func main() {
	// _, err := repository.NewGormDB()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// customizer := server.NewCustomizer()
	// server := srv.NewServer(customizer)
	// server.Start()

	container, err := di.InitializeApplication()
	if err != nil {
		fmt.Println(err)
		return
	}
	container.Server.Start()
}
