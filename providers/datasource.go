package providers

import (
	"github.com/bbcloudGroup/gothic/di"
	"gothic_app/datasource"
)

func RegisterDatabase(container di.Container) {

	container.Register(datasource.NewMovies)

	container.Register(datasource.NewAdmin)
	container.Register(datasource.NewCache)
}

