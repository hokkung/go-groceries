package main

import (
	"github.com/hokkung/go-groceries/server"
	srv "github.com/hokkung/srv/server"
)

func main() {
	customizer := server.NewCustomizer()
	server := srv.NewServer(customizer)
	server.Start()
}
