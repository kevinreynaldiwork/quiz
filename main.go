package main

import (
	"Quiz/database"
	"Quiz/routers"
)

func main() {
	db := database.Connect()
	database.DBMigrate(db)

	// Setup routes
	r := routers.SetupRouter(db)

	// Run server
	r.Run(":8080")
}
