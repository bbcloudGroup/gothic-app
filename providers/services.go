package providers

import (
	"github.com/bbcloudGroup/gothic/di"
	"gothic-app/business/services"
)

func RegisterServices(container di.Container) {

	container.Register(services.NewMovieService)

}
