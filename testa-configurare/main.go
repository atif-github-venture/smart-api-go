package main

import (
	"smartapigo/testa-configurare/app"
	"smartapigo/testa-configurare/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}

//italian-latin - testa-configurare -> test configuration
