package main

import (
	"smartapigo/action-list/app"
	"smartapigo/action-list/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}
