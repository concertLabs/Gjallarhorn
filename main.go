package main

import (
	"flag"

	"github.com/quiteawful/Gjallarhorn/importer"
	"github.com/quiteawful/Gjallarhorn/web"
)

var (
	WebApp   web.WebApp
	Importer importer.Importer

	configfile *string = flag.String("config", "config.json", "the config file to use")
)

func main() {
	flag.Parse()
	c, err := loadConfig(*configfile)
	if err != nil {
		panic(err)
	}
	Importer = importer.NewImporter(c.Importer.ScanDir)
	WebApp = web.NewWebApp(c.Httpd.RootDir)

	//go Importer.Run()
	WebApp.Run()
}
