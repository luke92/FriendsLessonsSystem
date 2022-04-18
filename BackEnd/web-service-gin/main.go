package main

import (
	"log"
	"os"

	di "github.com/luke92/FriendsLessonsSystem/di"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server, diErr := di.InitializeAPI()
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		log.Println("listening on", port)
		server.Start(port)
	}
}
