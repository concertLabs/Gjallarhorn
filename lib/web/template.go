package web

import (
	"html/template"
	"io"
	"log"
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

func (p *Renderer) loadTemplate(name, file string) error {
	if !strings.HasSuffix(file, ".html") {
		file += ".html"
	}
	basefile := path.Join(p.AssetDir, "templates", "main.html")
	tempfile := path.Join(p.AssetDir, "templates", file)

	// Maybe we can cache this
	var err error
	p.T = template.New(name)
	p.T, err = p.T.ParseFiles(basefile, tempfile)
	if err != nil {
		return err
	}

	return nil
}

func (p *Renderer) Render(name, file string, w io.Writer, data interface{}) error {
	err := p.loadTemplate(name, file)
	if err != nil {
		log.Printf("error while loading template %s: %v\n", file, err)
		return err
	}
	err = p.T.ExecuteTemplate(w, name, data)
	// err = p.T.Execute(w, data)
	if err != nil {
		log.Printf("error while rendering %s: %v\n", p.T.Name(), err)
	}
	return nil
}
