package controllers

import(
	// "github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)
type ExampleController struct {}

func (c *ExampleController) Get() mvc.Result  {
	return mvc.Response{
		ContentType:"text/html",
		Text:"<h1>Welcome</h1>",
	}
}

func (c *ExampleController) GetPing() string  {
	return "pong"
}

func (c *ExampleController) GetHello() interface{}  {
	return map[string]string{"message":"Hello Iris!"}
}