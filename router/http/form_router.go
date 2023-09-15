package http

import (
	"github.com/ChaukeBasani/gofarm/controller/context/pages"
	"github.com/labstack/echo/v4"
)

func FormRouter(app *echo.Echo) {
	app.POST("/record", pages.FormContext)
}