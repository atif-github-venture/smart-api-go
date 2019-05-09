package main

import (
	"smartapigo/data-setup/app"
	"smartapigo/data-setup/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}
