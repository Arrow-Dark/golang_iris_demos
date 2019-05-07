package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"golang_iris_demos/iris_demo_01/controllers"
)

func main() {
	app := iris.New()
	// app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	mvc.New(app).Handle(new(controllers.ExampleController))

	app.Run(iris.Addr(":8080"))
	
}

