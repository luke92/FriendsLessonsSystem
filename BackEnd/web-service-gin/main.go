package main

import (
	"log"
	di "github.com/luke92/FriendsLessonsSystem/di"
)

func main() {
	server, diErr := di.InitializeAPI()
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
