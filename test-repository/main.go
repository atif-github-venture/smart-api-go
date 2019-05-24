package main

import (
	"smartapigo/test-repository/app"
	"smartapigo/test-repository/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8081")
}
