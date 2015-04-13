package main

import (
	"flag"

	"github.com/quiteawful/Gjallarhorn/importer"
	"github.com/quiteawful/Gjallarhorn/web"
)

var (
	WebApp   web.WebApp
	Importer importer.Importer

	cfgHttpRoot  = flag.String("httproot", "./web/html/", "Root directory of the webhandler")
	cfgImportDir = flag.String("importdir", "./data", "basic data input directory.")
)

func main() {
	flag.Parse()

	Importer = importer.NewImporter(*cfgImportDir)
	WebApp = web.NewWebApp(*cfgHttpRoot)

	go Importer.Run()
	WebApp.Run()
}
