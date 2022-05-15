package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.New()

	app.Handle("GET", "/userinfo", func(context *context.Context) {
		path := context.Path()
		app.Logger().Info(path)
	})

	app.Get("/hello/{name}", func(context *context.Context) {
		path := context.Path()
		app.Logger().Info(path)
	})

	app.Get("/api/users/{userid:uint64}", func(context *context.Context) {
		userId, err := context.Params().GetUint64("userid")
		if err != nil {
			context.JSON(map[string]interface{}{
				"requestcode": 201,
				"message":     "bad request",
			})
			return
		}

		context.JSON(map[string]interface{}{
			"requestcode": 200,
			"userid":      userId,
		})
		//app.Logger().Info(panic)
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
