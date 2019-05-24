package main

import (
	"smartapigo/archive/makros-testaos-testa/app"
	"smartapigo/archive/makros-testaos-testa/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}

//french-latin - detail-testa -> detail test
