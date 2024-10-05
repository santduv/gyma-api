package main

import (
	"github.com/santduv/gyma-api/cmd"
	"github.com/santduv/gyma-api/internal/database"
)

func main() {
	app := cmd.CreateApp()
	database.ConnectToMongo()

	app.Listen(":3000")
}
