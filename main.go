package main

import (
	"github.com/ChaukeBasani/gofarm/constant"
	"github.com/ChaukeBasani/gofarm/router"
	"github.com/ChaukeBasani/gofarm/server"
	"github.com/labstack/echo/v4"
)

func main() {

	app := echo.New()
	//load static
	constant.LoadStatic(app)
	//load template folder
	app.Renderer = constant.LoadTemplate()
	//router
	router.LoadAllRouter(app)
	//server
	server.SetServer(app)
}
