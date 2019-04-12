package main

import (
	"smartapigo/mikros-testa/app"
	"smartapigo/mikros-testa/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}

//italian-latin - mikros-testa -> micro test
