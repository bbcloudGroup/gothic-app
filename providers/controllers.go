package providers

import (
	"github.com/bbcloudGroup/gothic/di"
	"gothic_app/web/controllers"
)

func RegisterController(container di.Container) {

	container.Register(controllers.NewHelloController)
	container.Register(controllers.NewMovieController)

}
