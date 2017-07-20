package main

import (
	"flag"
	"log"

	"github.com/quiteawful/Gjallarhorn/web"
)

func main() {
	configfile := flag.String("config", "config.json", "the json formatted configuration file")
	flag.Parse()
	c, err := loadConfig(*configfile)
	if err != nil {
		log.Fatalf("could not load configfile: %v\n", err)
	}

	// Importer := importer.NewImporter(c.Importer.ScanDir)
	WebApp := web.NewWebApp(c.Httpd.RootDir)

	//go Importer.Run()
	WebApp.Run()
}
