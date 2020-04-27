package providers

import (
	"github.com/bbcloudGroup/gothic/di"
	"gothic-app/boot/console"
)

func RegisterConsole(container di.Container) {

	container.Register(console.NewPrinter)

}
