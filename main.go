package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"

	"github.com/srueg/cala-iris/controllers"
)

func main() {
	app := iris.New()

	app.Use(recover.New())

	app.Use(logger.New())

	app.Logger().SetLevel("debug")

	app.RegisterView(
		iris.Handlebars("./templates", ".hbs").
			Layout("layout.hbs").
			Reload(true))

	app.StaticWeb("/public", "./public")

	mvc.New(app).Handle(new(controllers.HomeController))

	mvc.New(app.Party("/profile")).Handle(new(controllers.ProfileController))

	app.Run(iris.Addr("127.0.0.1:8080"), iris.WithoutServerError(iris.ErrServerClosed))

}
