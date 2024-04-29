package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	e.Static("/", "public/assets")
	e.GET("/", IndexHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

func IndexHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "template", "World")
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
