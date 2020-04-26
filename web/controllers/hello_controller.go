package controllers

import (
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"gothic_app/datasource"
)


type HelloController struct{
	admin datasource.Admin
	cache datasource.Cache
}

func NewHelloController(admin datasource.Admin, cache datasource.Cache) HelloController {
	return HelloController{admin:admin, cache:cache}
}

//var helloView = mvc.View{
//	Name: "gothic_app/index.html",
//	Data: map[string]interface{}{
//		"Title":     "Hello Page",
//		"MyMessage": "Welcome to my awesome website",
//	},
//}


func (c *HelloController) Get(ctx context.Context) mvc.Result {

	mes := make(chan string)

	go func() {
		i, err := c.cache.Get("abc").Result()
		//val, err := c.cache.Set("abc", 1, 0).Result()
		if err != nil {
			mes <- err.Error()
			return
		}

		mes <- i

	}()

	msg := <-mes


	if c.admin.HasTable(&datasource.User{}) {
		return mvc.Response{
			Text: "ddddddddddd",
			Code: 200,
		}
	}

	return mvc.Response{
		Text: msg,
		Code: 200,
	}
}


//
//// you can define a standard error in order to re-use anywhere in your app.
//var errBadName = errors.New("bad name")
//
//// you can just return it as error or even better
//// wrap this error with an mvc.Response to make it an mvc.Result compatible type.
//var badName = mvc.Response{Err: errBadName, Code: 400}
//
//// GetBy returns a "Hello {name}" response.
//// Demos:
//// curl -i http://localhost:8080/gothic_app/iris
//// curl -i http://localhost:8080/gothic_app/anything
//func (c *HelloController) GetBy(name string) mvc.Result {
//	if name != "iris" {
//		return badName
//		// or
//		// GetBy(name string) (mvc.Result, error) {
//		//	return nil, errBadName
//		// }
//	}
//
//	// return mvc.Response{Text: "Hello " + name} OR:
//	return mvc.View{
//		Name: "gothic_app/name.html",
//		Data: name,
//	}
//}
