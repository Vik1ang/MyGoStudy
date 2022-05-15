package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main1() {
	app := iris.New()

	userParty := app.Party("/users", func(context *context.Context) {
		context.Next()
	})

	userParty.Get("/register", func(context *context.Context) {
		app.Logger().Info("用户注册")
		context.HTML("<h1>Register</h1>")
	})

	usersRouter := app.Party("/admin", userMiddleware)

	usersRouter.Done(func(context *context.Context) {
		context.Application().Logger().Infof("response sent to " + context.Path())
	})

	usersRouter.Get("/info", func(context *context.Context) {
		context.HTML("<h1>info</h1>")
		context.Next()
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

func userMiddleware(context iris.Context) {
	context.Next()
}
