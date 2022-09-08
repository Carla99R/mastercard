package main

import (
	"github.com/kataras/iris/v12"
	"masterCard/utils/signature"
)

func main() {

	app := iris.New()
	app.Post("/example", signature.Sign)
	//os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
	//ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	//	"basePath": "v2",
	//})
	app.Listen(":8080")
}
