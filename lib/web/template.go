package web

import (
	"html/template"
	"path"
	"strings"
)

type Renderer struct {
	Name     string
	File     string
	AssetDir string
	T        *template.Template
}

func NewRenderer(assetDir string) *Renderer {
	return &Renderer{AssetDir: assetDir}
}

func (p *Renderer) LoadTemplate(name, file string) (*template.Template, error) {
	if !strings.HasSuffix(file, ".html") {
		file += ".html"
	}
	basefile := path.Join(p.AssetDir, "templates", "_base.html")
	tempfile := path.Join(p.AssetDir, "templates", file)

	t := template.New(name)
	t, err := t.ParseFiles(basefile, tempfile)
	if err != nil {
		return nil, err
	}

	return t, nil
}
