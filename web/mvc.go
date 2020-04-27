package web

import (
	"github.com/bbcloudGroup/gothic/di"
	"github.com/bbcloudGroup/gothic/routes"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"gothic-app/web/controllers"
	"gothic-app/web/middleware"
)


var middlewares = [] context.Handler {
	middleware.K2,
	middleware.K,
}

func Mvc(app *iris.Application) map[string]routes.Route {

	return map[string]routes.Route {
		"/hello": routes.New(Hello),//.WithMiddleware(middleware.K).WithTerminate(middlewares...),
		"/movies": routes.New(Movie).WithMiddleware(middleware.BasicAuth),
	}

}

func Hello(m *mvc.Application) {
	di.Invoke(func(controller controllers.HelloController) {m.Handle(&controller)})
	//m.Handle(new(controllers.HelloController))
}

func Movie(m *mvc.Application) {
	di.Invoke(func(controller controllers.MovieController) {m.Handle(&controller)})
	//repo := repositories.NewMovieRepository(database.Movies)
	//movieService := services.NewMovieService(repo)
	//m.Register(movieService)
	//m.Handle(new(controllers.MovieController))
}

