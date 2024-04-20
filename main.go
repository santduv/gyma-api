package main

import (
	"github.com/santduv/gyma-api/cmd"
)

func main() {
	app := cmd.CreateApp()

	app.Listen(":3000")
}
