package main

import (
	"smartapigo/agere-lystan/app"
	"smartapigo/agere-lystan/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}

//latin-german - agere-lystan -> action list test
