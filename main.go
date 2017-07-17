package main

import (
	"flag"

	"github.com/quiteawful/Gjallarhorn/importer"
	"github.com/quiteawful/Gjallarhorn/web"
)

func main() {
	configfile := flag.String("config", "config.json", "the json formatted configuration file")
	flag.Parse()
	c, err := loadConfig(*configfile)
	if err != nil {
		panic(err)
	}

	Importer := importer.NewImporter(c.Importer.ScanDir)
	WebApp := web.NewWebApp(c.Httpd.RootDir)

	go Importer.Run()
	WebApp.Run()
}
