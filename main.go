package main

import (
	"github.com/bbcloudGroup/gothic/bootstrap"
	"gothic_app/boot"
)


func main() {

	app := boot.NewApp(bootstrap.GetArgs())
	app.Run()

}
