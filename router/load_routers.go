package router

import (
	"github.com/ChaukeBasani/gofarm/router/http"
	"github.com/labstack/echo/v4"
)

func LoadAllRouter(app *echo.Echo) {
	http.IndexRouter(app)
}
