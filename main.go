package main

import (
	dotenv "github.com/joho/godotenv"
	"github.com/qcodelabsllc/qreeket/gateway/network"
)

func main() {
	// load env
	if err := dotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	// start server
	network.InitServer()
}
