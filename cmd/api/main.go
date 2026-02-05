package main

import "github.com/JoaoVitor615/URL-shortener/internal/server"

func main() {
	deps := server.NewDependencies()
	server.Run(deps)
}
