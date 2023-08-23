package main

import (
	"fmt"

	"github.com/hokkung/go-groceries/repository"
	"github.com/hokkung/go-groceries/server"
	srv "github.com/hokkung/srv/server"
)

func main() {
	_, err := repository.NewGormDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	customizer := server.NewCustomizer()
	server := srv.NewServer(customizer)
	server.Start()
}
