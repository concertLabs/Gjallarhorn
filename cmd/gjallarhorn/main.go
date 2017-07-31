package main

import (
	"flag"
	"log"

	"github.com/quiteawful/Gjallarhorn/lib/config"
	"github.com/quiteawful/Gjallarhorn/lib/importer"
	"github.com/quiteawful/Gjallarhorn/web"
)

// Args is a struct to parse the commandline arguments
type Args struct {
	config string
}

// DefaultOptions returns a struct with the minimal/default values
// might be extendet later on
func defaultArgs() Args {
	return Args{config: "config.json"}
}

func parseArgs() Args {
	var result = defaultArgs()

	flag.StringVar(&result.config, "config", result.config, "the json formatted config file")
	flag.Parse()

	return result
}

func main() {
	args := parseArgs()
	c, err := config.Open(args.config)
	if err != nil {
		log.Fatalf("could not open configfile: %v\n", err)
	}

	Importer := importer.NewImporter(c.Importer)
	WebApp := web.NewWebApp(c.Httpd.RootDir)

	go Importer.Run()
	WebApp.Run()
}
