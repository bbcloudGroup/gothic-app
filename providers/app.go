package providers

import "github.com/bbcloudGroup/gothic/di"

func Register(container di.Container) {
	RegisterDatabase(container)
	RegisterRepositories(container)
	RegisterServices(container)
	RegisterController(container)
	RegisterConsole(container)
}