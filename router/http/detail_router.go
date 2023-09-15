package http

import (
	"github.com/ChaukeBasani/gofarm/controller/context/pages"
	"github.com/labstack/echo/v4"
)

func DetailsRouter(app *echo.Echo) {
	app.GET("/:productId", pages.DetailsContext)
}