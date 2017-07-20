package web

import (
	"html/template"
	"path"
	"strings"
)

func (app *WebApp) loadTemplate(name, file string) (*template.Template, error) {
	if !strings.HasSuffix(file, ".html") {
		file += ".html"
	}
	basefile := path.Join(app.RootDir, "templates", "_base.html")
	tempfile := path.Join(app.RootDir, "templates", file)

	t := template.New(name)
	t, err := t.ParseFiles(basefile, tempfile)
	if err != nil {
		return nil, err
	}
	return t, nil
}
