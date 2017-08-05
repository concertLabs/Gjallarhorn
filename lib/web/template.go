package web

import (
	"html/template"
	"path"
	"strings"
)

func (app *App) loadTemplate(name, file string) (*template.Template, error) {
	if !strings.HasSuffix(file, ".html") {
		file += ".html"
	}
	basefile := path.Join(app.AssetDir, "templates", "_base.html")
	tempfile := path.Join(app.AssetDir, "templates", file)

	t := template.New(name)
	t, err := t.ParseFiles(basefile, tempfile)
	if err != nil {
		return nil, err
	}
	return t, nil
}
