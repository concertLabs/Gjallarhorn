package main

import (
	"github.com/quiteawful/Gjallarhorn/importer"
	"github.com/quiteawful/Gjallarhorn/web"
)

var (
	WebApp   web.WebApp
	Importer importer.Importer
)

func main() {
	c := loadConfig()

	Importer = importer.NewImporter(c.ImportDirectory)
	WebApp = web.NewWebApp(c.HttpRoot)

	go Importer.Run()
	WebApp.Run()
}
