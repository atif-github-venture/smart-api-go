package main

import (
	"smartapigo/create-project/app"
	"smartapigo/create-project/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}
