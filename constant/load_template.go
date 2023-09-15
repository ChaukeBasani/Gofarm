package constant

import (
	"fmt"
	"html/template"
	"io"

	"path/filepath"

	"os"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func LoadTemplate() *Template {

	path, _ := os.Executable()
	filepath := filepath.Dir(path)
	fmt.Print(filepath)
	templateFolder := fmt.Sprintf("%v/repository/templates/*", filepath)

	template := &Template{
		templates: template.Must(template.ParseGlob(templateFolder)),
	}

	return template
}
