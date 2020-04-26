package errors

import (
	"github.com/bbcloudGroup/gothic/logerr"
	"github.com/kataras/iris/v12"
)

func Handler(ctx iris.Context) {

	switch ctx.GetStatusCode() {
	default:
		_, err := ctx.Problem(logerr.NewProblem(ctx))
		if err != nil {
			panic(err)
		}
		break
	}
}
