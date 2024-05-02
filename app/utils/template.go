package utils

import (
	"html/template"
	"io"
	"io/fs"

	"github.com/labstack/echo"
)

type Template struct {
	FS fs.FS
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return template.Must(template.ParseFS(
		t.FS,
		"web/views/layouts/*.html",
		"web/views/"+name+".html",
	)).Execute(w, data)
}

// return template.Must(template.ParseFS(
// 	WebFS,
// 	"web/views/layouts/*.html",
// 	"web/views/"+name+".html",
// )).Execute(w, data)
