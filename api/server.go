package api

import (
	"embed"
	"html/template"
	"io"

	"github.com/labstack/echo"
)

//go:embed web/views/*.html
var tmplFS embed.FS

func Setup() {
	t := &Template{
		Templates: template.Must(template.ParseFS(tmplFS, "web/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	SetupRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl := template.Must(t.Templates.Clone())
	tmpl = template.Must(tmpl.ParseFS(tmplFS, "web/views/"+name))
	return tmpl.ExecuteTemplate(w, name, data)
}
