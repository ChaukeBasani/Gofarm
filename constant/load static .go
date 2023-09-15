package constant

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func LoadStatic(app *echo.Echo) {
	path, _ := os.Executable()
	filepath := filepath.Dir(path)
	fmt.Print(filepath)
	staticFolder := fmt.Sprintf("%v/repository/assets", filepath)
	// load static path.. utilities to make our profile to look goo
	// they nomarlly doeas mot change
	app.Static("static", staticFolder)

}
