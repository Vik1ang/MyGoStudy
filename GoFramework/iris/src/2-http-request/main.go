package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	app := iris.New()

	// GET
	app.Get("/getRequest", func(context *context.Context) {
		path := context.Path()
		app.Logger().Info(path)
	})

	app.Get("/userpath", func(context *context.Context) {
		path := context.Path()
		app.Logger().Info(path)

		context.WriteString("请求路径:" + path)
	})

	app.Get("/userinfo", func(context *context.Context) {
		path := context.Path()
		app.Logger().Info(path)

		userName := context.URLParam("username")
		app.Logger().Info(userName)

		pwd := context.URLParam("pwd")
		app.Logger().Info(pwd)

		context.HTML("<h1>" + userName + "," + pwd + "</h1>")
	})

	// POST
	app.Post("/postLogin", func(context *context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		name := context.PostValue("name")
		pwd := context.PostValue("pwd")
		app.Logger().Info(name, " ", pwd)
		context.HTML(name)
	})

	app.Post("/postJson", func(context *context.Context) {
		path := context.Path()
		app.Logger().Info("URL:", path)

		var person Person
		if err := context.ReadJSON(&person); err != nil {
			panic(err.Error())
		}

		context.Writef("Received: %#+v\n", person)
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

}
