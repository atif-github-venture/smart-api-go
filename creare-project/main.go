package main

import (
	"smartapigo/creare-project/app"
	"smartapigo/creare-project/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}

//latin - creare -> create
