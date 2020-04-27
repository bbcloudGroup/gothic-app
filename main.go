package main

import (
	"github.com/bbcloudGroup/gothic/bootstrap"
	"gothic-app/boot"
)


func main() {

	app := boot.NewApp(bootstrap.GetArgs())
	app.Run()

}
