package env

import (
	"fmt"
	"github.com/joho/godotenv"
)

func MustLoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("cannot find .env file")
	}
}
