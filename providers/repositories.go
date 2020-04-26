package providers

import (
	"github.com/bbcloudGroup/gothic/di"
	"gothic_app/business/repositories"
)

func RegisterRepositories(container di.Container) {

	container.Register(repositories.NewMovieRepository)

}