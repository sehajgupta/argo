package main

import (
	"trippin/db"
	"trippin/routes"
)

func main() {
	// initialize the database connection
	db.Init()
	defer db.DB.Close()

	// setup routes
	routes.SetupRouter().Run(":8080")
}