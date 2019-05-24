package main

import (
	"smartapigo/test-config/app"
	"smartapigo/test-config/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}
