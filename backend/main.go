package main

import (
	"trippin/internal"
	"trippin/db"
)

func main() {
	// initialize the database connection
	db.Init()
	defer db.DB.Close()

	// init server
	server := server.NewServer(db.DB)

	server.SetupRouterGroup(server.Router)
	// serve routes
	server.Router.Run(":8080")
}