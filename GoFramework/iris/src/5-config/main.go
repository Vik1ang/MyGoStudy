package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Configure(iris.WithConfiguration(iris.Configuration{
		// 如何设置为true, 人为中断程序执行时, 不会自动正常关闭服务器, 需要自定义处理
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  true,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon,02 Jan 2006 15:04:05 GMT",
		Charset:                           "utf-8",
	}))

	app.Configure(iris.WithConfiguration(iris.TOML("./config/iris.toml")))

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
