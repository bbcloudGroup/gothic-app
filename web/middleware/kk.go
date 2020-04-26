package middleware

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func K(ctx iris.Context) {

	fmt.Print(ctx.Request().RequestURI)
	ctx.Next()
}


func K2(ctx iris.Context) {
	ctx.WriteString("\nFrom Done Handler")

	//fmt.Print("aab")
}
