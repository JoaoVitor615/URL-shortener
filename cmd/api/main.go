package main

import (
	"log"

	"github.com/JoaoVitor615/URL-shortener/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Warning: .env file not found, using system environment variables.")
	}

	deps := server.NewDependencies()
	server.Run(deps)
}
