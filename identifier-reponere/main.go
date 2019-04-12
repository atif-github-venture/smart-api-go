package main

import (
	"smartapigo/identifier-reponere/app"
	"smartapigo/identifier-reponere/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}


//latin - repository -> reponere