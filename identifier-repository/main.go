package main

import (
	"smartapigo/identifier-repository/app"
	"smartapigo/identifier-repository/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}
