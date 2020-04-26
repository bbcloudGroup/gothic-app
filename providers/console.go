package providers

import (
	"github.com/bbcloudGroup/gothic/di"
	"gothic_app/boot/console"
)

func RegisterConsole(container di.Container) {

	container.Register(console.NewPrinter)

}
