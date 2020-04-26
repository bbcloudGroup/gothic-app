package boot

import (
	"github.com/bbcloudGroup/gothic/bootstrap"
	"github.com/bbcloudGroup/gothic/cli"
	"github.com/bbcloudGroup/gothic/config"
	"github.com/bbcloudGroup/gothic/database"
	"github.com/bbcloudGroup/gothic/di"
	"github.com/bbcloudGroup/gothic/logerr"
	"github.com/bbcloudGroup/gothic/routes"
	"github.com/bbcloudGroup/gothic/static"
	"gothic_app/boot/console"
	"gothic_app/boot/errors"
	"gothic_app/boot/properties"
	"gothic_app/providers"
	"gothic_app/web"
)


func NewApp(args bootstrap.Args) *bootstrap.Bootstrapper {

	app := bootstrap.New()
	app.SetupUp(
		config.BootStrap(args.Env, properties.BootStrap),
		static.BootStrap(),
		logerr.BootStrap(errors.Handler),
		database.BootStrap(),
		di.BootStrap(providers.Register),
		routes.BootStrap(web.Mvc),
		cli.BootStrap(console.Schedule),
		//iris.WithoutServerError(iris.ErrServerClosed),
		//iris.WithOptimizations,
		//iris.WithoutBanner,
		// ...
		)




	//app.Handle("GET", "/", func(ctx iris.Context) {
	//	//ctx.Application().Logger().Infof("hello: %s", "i am info.")
	//	s := config.GetString(ctx.Application(), "Author")
	//	ctx.HTML(s)
	//})

	return app
}






//package bootstrap
//
//import (
//	"github.com/bbcloudGroup/gothic/		ctx.Application().R
//	"github.com/kataras/iris/v12"
//	"github.com/kataras/iris/v12/middleware/logger"
//	"github.com/kataras/iris/v12/middleware/recover"
//	"github.com/kataras/iris/v12/mvc"
//)
//
//func NewApp() *iris.Application {
//	app := iris.New()
//	app.Use(recover.New())
//	app.Use(logger.New())
//
//
//	// Serve a controller based on the root Router, "/".
//	mvc.New(app).Handle(new(ExampleController))
//	return app
//}
//
//
//func Http(app *iris.Application) {
//
//	var m map[string]interface{}    //声明变量，不分配内存
//	m = make(map[string]interface{}) //必可不少，分配内存
//	m["name"] = "simon"
//	var age int = 12
//	m["age"] = age
//	m["addr"] = "China"
//
//	x := iris.YAML("config/iris.yaml")
//	x.Other = m
//
//	config.Init()
//
//	_ = app.Run(iris.Addr(":8080"), iris.WithConfiguration(x))
//}
//
//
//
//// ExampleController serves the "/", "/ping" and "/hello".
//type ExampleController struct{}
//
//// Get serves
//// Method:   GET
//// Resource: http://localhost:8080
//func (c *ExampleController) Get() mvc.Result {
//	return mvc.Response{
//		ContentType: "text/html",
//		Text:        "<h1>Welcome</h1>",
//	}
//}
//
//// GetPing serves
//// Method:   GET
//// Resource: http://localhost:8080/ping
//func (c *ExampleController) GetPing(ctx iris.Context) string {
//
//	config.Get(ctx.Application(), "name")
//
//	return ctx.Application().ConfigurationReadOnly().GetOther()["name"].(string)
//}
//
//// GetHello serves
//// Method:   GET
//// Resource: http://localhost:8080/hello
//func (c *ExampleController) GetHello() interface{} {
//	return map[string]string{"message": "Hello Iris!"}
//}
//
//// BeforeActivation called once, before the controller adapted to the main application
//// and of course before the server ran.
//// After version 9 you can also add custom routes for a specific controller's methods.
//// Here you can register custom method's handlers
//// use the standard router with `ca.Router` to do something that you can do without mvc as well,
//// and add dependencies that will be binded to a controller's fields or method function's input arguments.
//func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
//	anyMiddlewareHere := func(ctx iris.Context) {
//		ctx.Application().Logger().Warnf("Inside /custom_path")
//		ctx.Next()
//	}
//	b.Handle("GET", "/custom_path", "CustomHandlerWithoutFollowingTheNamingGuide", anyMiddlewareHere)
//
//	// or even add a global middleware based on this controller's router,
//	// which in this example is the root "/":
//	// b.Router().Use(myMiddleware)
//}
//
//// CustomHandlerWithoutFollowingTheNamingGuide serves
//// Method:   GET
//// Resource: http://localhost:8080/custom_path
//func (c *ExampleController) CustomHandlerWithoutFollowingTheNamingGuide() string {
//	return "hello from the custom handler without following the naming guide"
//}

